package com.example.testapp.fragments;

import android.os.Bundle;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.TextView;

import androidx.fragment.app.Fragment;
import androidx.annotation.NonNull;
import androidx.annotation.Nullable;
import androidx.recyclerview.widget.LinearLayoutManager;
import androidx.recyclerview.widget.RecyclerView;

import com.example.testapp.App;
import com.example.testapp.adapters.CommentsAdapter;
import com.example.testapp.greendao.DaoSession;
import com.example.testapp.greendao.StationDao;
import com.example.testapp.utils.Helper;
import com.example.testapp.R;
import com.example.testapp.models.Station;
import com.google.android.material.floatingactionbutton.FloatingActionButton;

import org.greenrobot.greendao.query.Query;

import java.util.concurrent.TimeUnit;

import io.reactivex.android.schedulers.AndroidSchedulers;
import io.reactivex.disposables.CompositeDisposable;
import io.reactivex.observers.DisposableSingleObserver;
import io.reactivex.schedulers.Schedulers;
import retrofit2.Response;

public class InfoFragment extends Fragment {

    private CompositeDisposable disposable = new CompositeDisposable();
    private App app;
    private RecyclerView recyclerView;
    private CommentsAdapter commentsAdapter;
    private Station currentStation;
    private FloatingActionButton buttonEditStation;
    private FloatingActionButton buttonCreateComment;

    public InfoFragment(double lat, double lng) {
        this.currentStation = new Station();
        this.currentStation.setLatitude(lat);
        this.currentStation.setLongitude(lng);
    }

    @Nullable
    @Override
    public View onCreateView(@NonNull LayoutInflater inflater,
                             @Nullable ViewGroup container,
                             @Nullable Bundle savedInstanceState) {
        View view = inflater.inflate(R.layout.fragment_info, container, false);
        return view;
    }

    @Override
    public void onViewCreated(@NonNull View view, @Nullable Bundle savedInstanceState) {
        super.onViewCreated(view, savedInstanceState);
        this.app = (App) getActivity().getApplication();
        buttonEditStation = (FloatingActionButton) getView().findViewById(R.id.floatingActionButtonEditStation);
        buttonEditStation.setOnClickListener(this::onButtonEditStation);
        buttonCreateComment = (FloatingActionButton) getView().findViewById(R.id.floatingActionButtonCreateComment);
        buttonCreateComment.setOnClickListener(this::onButtonCreateComment);
        updateInfo();
    }

    public void updateInfo() {
        disposable.add(app.getElchargeService().getStationApi().readStationsByLatAndLng(0, 0, currentStation.getLatitude(), currentStation.getLongitude())
                .subscribeOn(Schedulers.io())
                .timeout(3000, TimeUnit.MILLISECONDS)
                .observeOn(AndroidSchedulers.mainThread())
                .subscribeWith(new StationDisposableSingleObserver(this)));
    }

    private void updateInfoOnView() {
        if (app.getElchargeService().getUser().getId().equals(currentStation.getOwnerID())) {
            buttonEditStation.show();
        } else {
            buttonEditStation.hide();
        }
        recyclerView = getView().findViewById(R.id.recyclerViewComments);
        recyclerView.setLayoutManager((new LinearLayoutManager(App.getAppContext())));

        commentsAdapter = new CommentsAdapter(currentStation.getComments());
        recyclerView.setAdapter(commentsAdapter);

        ((TextView) getView().findViewById(R.id.stationId)).setText(currentStation.getId());
        ((TextView) getView().findViewById(R.id.stationName)).setText(currentStation.getStationName());
        ((TextView) getView().findViewById(R.id.rating)).setText(String.format("%.2f", currentStation.getRating()) + " / 5");
        ((TextView) getView().findViewById(R.id.description)).setText(currentStation.getDescription());
        ((TextView) getView().findViewById(R.id.date)).setText(Helper.getDateFromISO8601(currentStation.getUpdateAt()));
    }

    private void onButtonEditStation(View v) {
        getFragmentManager().beginTransaction().add(R.id.container, new EditStationFragment(currentStation)).commit();
    }

    private void onButtonCreateComment(View v) {
        getFragmentManager().beginTransaction().add(R.id.container, new CreateCommentFragment(currentStation.getId())).commit();
    }

    private class StationDisposableSingleObserver extends DisposableSingleObserver<Response<Station>> {
        private InfoFragment parent;

        public StationDisposableSingleObserver(InfoFragment parent) {
            this.parent = parent;
        }

        @Override
        public void onSuccess(Response<Station> response) {
            try {
                if (response.code() == 200) {
                    Station station = response.body();
                    Helper.messageLogger(App.getAppContext(), Helper.LogType.INFO, "station", Integer.toString(response.code()));
                    System.out.println(station.toString());
                    currentStation = station;
                    updateInfoOnView();
                } else {
                    Helper.messageLogger(App.getAppContext(), Helper.LogType.INFO, "station", Integer.toString(response.code()));
                    if (response.code() == 401) {
                        LoginFragment lf = new LoginFragment();
                        getFragmentManager().beginTransaction().add(R.id.container, lf).commit();
                        getFragmentManager().beginTransaction().show(lf).commit();
                    } else {
                        getFragmentManager().beginTransaction().remove(parent).commit();
                    }
                }

            } catch (Exception e) {
                Helper.messageLogger(App.getAppContext(), Helper.LogType.ERR, "station", e.getMessage());
            }
        }

        @Override
        public void onError(Throwable e) {
            DaoSession daoSession = ((App) getActivity().getApplication()).getDaoSession();
            Station station = daoSession.queryBuilder(Station.class)
                    .where(StationDao.Properties.Latitude.eq(currentStation.getLatitude())
                            , (StationDao.Properties.Longitude.eq(currentStation.getLongitude()))
                    ).unique();
            Helper.messageLogger(App.getAppContext(), Helper.LogType.ERR, "station", "offline mode");
            currentStation = station;
            updateInfoOnView();
        }
    }

    @Override
    public void onDestroy() {
        disposable.dispose();
        super.onDestroy();
    }
}