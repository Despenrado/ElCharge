package com.example.testapp.fragments;

import android.os.Bundle;

import androidx.annotation.NonNull;
import androidx.fragment.app.ListFragment;

import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.ArrayAdapter;
import android.widget.Button;
import android.widget.ListView;

import com.example.testapp.App;
import com.example.testapp.greendao.DaoSession;
import com.example.testapp.models.Station;
import com.example.testapp.models.User;
import com.example.testapp.utils.Helper;
import com.example.testapp.R;

import java.util.List;

import io.reactivex.android.schedulers.AndroidSchedulers;
import io.reactivex.disposables.CompositeDisposable;
import io.reactivex.observers.DisposableSingleObserver;
import io.reactivex.schedulers.Schedulers;
import okhttp3.ResponseBody;
import retrofit2.Response;

public class SettingsFragment extends ListFragment {

    private CompositeDisposable disposable = new CompositeDisposable();
    private App app;

    public SettingsFragment() {
        // Required empty public constructor
    }

    private String[] data = new String[]{
            "sign in",
            "log out",
            "update offline data"
    };

    @Override
    public void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        ArrayAdapter<String> adapter;
        adapter = new ArrayAdapter<String>(App.getAppContext(),
                android.R.layout.simple_list_item_1,
                data);
        setListAdapter(adapter);
        app = (App) getActivity().getApplication();
    }

    @Override
    public View onCreateView(LayoutInflater inflater, ViewGroup container,
                             Bundle savedInstanceState) {
        return inflater.inflate(R.layout.fragment_settings, container, false);
    }

    @Override
    public void onListItemClick(@NonNull ListView l, @NonNull View v, int position, long id) {
        switch ((String) getListAdapter().getItem(position)) {
            case "sign in":
                getFragmentManager().beginTransaction().add(R.id.container, new LoginFragment()).commit();
                break;
            case "log out":
                System.out.println("logout");
                logout();
                break;
            case "update offline data":
                updateOfflineDatabase();
        }
    }

    // send request to backend: logout
    private void logout() {
        if (app.getElchargeService().getUser() != null) {
            disposable.add(app.getElchargeService().getUserApi().logout(app.getElchargeService().getUser().getId())
                    .subscribeOn(Schedulers.io())
                    .observeOn(AndroidSchedulers.mainThread())
                    .subscribeWith(new DisposableSingleObserver<ResponseBody>() {
                        @Override
                        public void onSuccess(ResponseBody responseBody) {
                            try {
                                if (responseBody != null) {
                                    app.getElchargeService().setToken(""); //remove token
                                    User tmpUser = new User();
                                    tmpUser.setId("");
                                    app.getElchargeService().setUser(tmpUser); //remove token
                                    LoginFragment lf = new LoginFragment();
                                    getFragmentManager().beginTransaction().add(R.id.container, lf).commit();
                                    getFragmentManager().beginTransaction().show(lf).commit();
                                }
                                Helper.messageLogger(App.getAppContext(), Helper.LogType.INFO, "logout", responseBody.string());
                            } catch (Exception e) {
                                Helper.messageLogger(App.getAppContext(), Helper.LogType.ERR, "logout", e.getMessage());
                            }
                        }

                        @Override
                        public void onError(Throwable e) {
                            Helper.messageLogger(App.getAppContext(), Helper.LogType.ERR, "logout", e.getMessage());
                        }
                    }));
        }
    }

    private void updateOfflineDatabase() {
        DaoSession daoSession = app.getDaoSession();
        List<Station> stationOfflineList = daoSession.loadAll(Station.class);
        for (Station station : stationOfflineList) {
            disposable.add(app.getElchargeService().getStationApi().findById(station.getId())
                    .subscribeOn(Schedulers.io())
                    .observeOn(AndroidSchedulers.mainThread())
                    .subscribeWith(new DisposableSingleObserver<Response<Station>>() {
                        @Override
                        public void onSuccess(Response<Station> response) {
                            if (response.code() == 302 || response.code() == 200) {
                                Station updatedStation = response.body();
                                daoSession.update(updatedStation);
                                return;
                            }
                            if (response.code() == 204) {
                                daoSession.delete(station);
                                return;
                            }
                            if (response.code() == 401) {
                                LoginFragment lf = new LoginFragment();
                                getFragmentManager().beginTransaction().add(R.id.container, lf).commit();
                                getFragmentManager().beginTransaction().show(lf).commit();
                                return;
                            }
                        }

                        @Override
                        public void onError(@NonNull Throwable e) {

                        }
                    }));
        }
    }

    @Override
    public void onDestroy() {
        disposable.dispose();
        super.onDestroy();
    }
}