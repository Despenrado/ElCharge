package com.example.testapp.fragments;

import android.os.Bundle;

import androidx.annotation.NonNull;
import androidx.annotation.Nullable;
import androidx.fragment.app.Fragment;

import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.Button;
import android.widget.EditText;
import android.widget.RatingBar;
import android.widget.TextView;

import com.example.testapp.App;
import com.example.testapp.Helper;
import com.example.testapp.R;
import com.example.testapp.api.models.Comment;
import com.example.testapp.api.models.Station;

import io.reactivex.android.schedulers.AndroidSchedulers;
import io.reactivex.disposables.CompositeDisposable;
import io.reactivex.observers.DisposableSingleObserver;
import io.reactivex.schedulers.Schedulers;
import retrofit2.Response;


public class EditStationFragment extends Fragment {

    private CompositeDisposable disposable = new CompositeDisposable();
    private App app;
    private Station station;
    private EditText editTextStationName;
    private EditText editTextDescription;

    public EditStationFragment(Station station) {
        this.station = station;
    }

    @Override
    public void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);

    }

    @Override
    public View onCreateView(LayoutInflater inflater, ViewGroup container,
                             Bundle savedInstanceState) {
        return inflater.inflate(R.layout.fragment_edit_station, container, false);
    }

    @Override
    public void onViewCreated(@NonNull View view, @Nullable Bundle savedInstanceState) {
        super.onViewCreated(view, savedInstanceState);
        this.app = (App) getActivity().getApplication();
        Button btnSave = (Button) getView().findViewById(R.id.buttonSaveEdit);
        btnSave.setOnClickListener(this::onButtonSave);
        Button btnCancel = (Button) getView().findViewById(R.id.buttonCancelEdit);
        btnCancel.setOnClickListener(this::onButtonCancel);
        ((TextView) getView().findViewById(R.id.textViewID)).setText(station.getId());
        ((TextView) getView().findViewById(R.id.textViewLat)).setText(Double.toString(station.getLatitude()));
        ((TextView) getView().findViewById(R.id.textViewLng)).setText(Double.toString(station.getLongitude()));
        editTextStationName = (EditText) getView().findViewById(R.id.editTextStationName);
        editTextStationName.setText(station.getStationName());
        editTextDescription = (EditText) getView().findViewById(R.id.editTextDescription);
        editTextDescription.setText(station.getDescription());
    }

    private void onButtonSave(View v) {
        station.setRating(0);
        station.setStationName(editTextStationName.getText().toString());
        station.setDescription(editTextDescription.getText().toString());
        saveStation();
    }

    private void onButtonCancel(View v) {
        getFragmentManager().beginTransaction().remove(this).commit();
    }

    private void saveStation() {
        final EditStationFragment tmp = this;
        disposable.add(app.getElchargeService().getStationApi().updateById(station.getId(), app.getElchargeService().getUser().getId(), station)
                .subscribeOn(Schedulers.io())
                .observeOn(AndroidSchedulers.mainThread())
                .subscribeWith(new DisposableSingleObserver<Response<Station>>() {
                    @Override
                    public void onSuccess(Response<Station> response) {
                        try {
                            if (response.code() == 201) {
                                Helper.messageLogger(App.getAppContext(), Helper.LogType.INFO, "station", "Saved!");
                                InfoFragment inf = (InfoFragment) getFragmentManager().findFragmentByTag("station_info");
                                if (inf != null) {
                                    inf.updateInfo();
                                }
                                getFragmentManager().beginTransaction().remove(tmp).commit();
                            } else {
                                Helper.messageLogger(App.getAppContext(), Helper.LogType.INFO, "station", response.message());
                                if (response.code() == 401) {
                                    LoginFragment lf = new LoginFragment();
                                    getFragmentManager().beginTransaction().add(R.id.container, lf).commit();
                                    getFragmentManager().beginTransaction().show(lf).commit();
                                }
                            }

                        } catch (Exception e) {
                            Helper.messageLogger(App.getAppContext(), Helper.LogType.ERR, "station", e.getMessage());
                        }
                    }

                    @Override
                    public void onError(Throwable e) {
                        Helper.messageLogger(App.getAppContext(), Helper.LogType.ERR, "station", e.getMessage());
                    }
                }));
    }
}