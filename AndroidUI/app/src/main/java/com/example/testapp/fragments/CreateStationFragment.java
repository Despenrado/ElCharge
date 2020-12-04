package com.example.testapp.fragments;

import android.os.Bundle;

import androidx.fragment.app.Fragment;

import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.Button;
import android.widget.EditText;

import com.example.testapp.App;
import com.example.testapp.Helper;
import com.example.testapp.R;
import com.example.testapp.api.models.Station;

import io.reactivex.android.schedulers.AndroidSchedulers;
import io.reactivex.disposables.CompositeDisposable;
import io.reactivex.observers.DisposableSingleObserver;
import io.reactivex.schedulers.Schedulers;
import retrofit2.Response;


public class CreateStationFragment extends Fragment {

    private CompositeDisposable disposable = new CompositeDisposable();
    private App app;

    EditText stationLocation;

    @Override
    public void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        app = (App) getActivity().getApplication();
    }

    @Override
    public View onCreateView(LayoutInflater inflater, ViewGroup container,
                             Bundle savedInstanceState) {
        View view = inflater.inflate(R.layout.fragment_create_station, container, false);
        Button btnCreate = (Button) view.findViewById(R.id.buttonCreateStation);
        btnCreate.setOnClickListener(this::onButtonSaveStationClick);
        Button btnMap = (Button) view.findViewById(R.id.buttonCheckOnMap);
        btnMap.setOnClickListener(this::onButtonViewMapClick);
        stationLocation = (EditText) view.findViewById(R.id.editTextLocation);
        return view;
    }

    public void onButtonSaveStationClick(View v) {
        String stationName = ((EditText) getView().findViewById(R.id.editTextStationName)).getText().toString();
        String stationLocation = this.stationLocation.getText().toString();
        String stationDescription = ((EditText) getView().findViewById(R.id.editTextDescription)).getText().toString();
        if (stationName.length() > 1 && stationLocation.length() > 2 && stationDescription.length() > 4) {
            String coordinates[] = stationLocation.split("[, |]");
            if (coordinates.length < 2) {
                Helper.messageLogger(App.getAppContext(), Helper.LogType.NONE, "station", "wrong location");
                return;
            }
            try {
                double lat = Double.parseDouble(coordinates[0]);
                double lng = Double.parseDouble(coordinates[1]);
                Station station = new Station();
                station.setStationName(stationName);
                station.setDescription(stationDescription);
                station.setLatitude(lat);
                station.setLongitude(lng);
                if (app.getElchargeService().getUser() != null) {
                    station.setOwnerID(app.getElchargeService().getUser().getId());
                }
                createStation(station);
            } catch (NumberFormatException e) {
                Helper.messageLogger(App.getAppContext(), Helper.LogType.ERR, "station", e.getMessage());
            }
        } else {
            Helper.messageLogger(App.getAppContext(), Helper.LogType.NONE, "station", "One of fields so short");
        }
    }

    // send request to backend: create new station
    public void createStation(Station station) {
        final CreateStationFragment tmp = this;
        disposable.add(app.getElchargeService().getStationApi().createStation(station)
                .subscribeOn(Schedulers.io())
                .observeOn(AndroidSchedulers.mainThread())
                .subscribeWith(new DisposableSingleObserver<Response<Station>>() {
                    @Override
                    public void onSuccess(Response<Station> response) {
                        try {
                            if (response.code() == 201) {
                                Helper.messageLogger(App.getAppContext(), Helper.LogType.INFO, "station", "Created successfully!");
                                getFragmentManager().beginTransaction().replace(R.id.container, new MapsFragment()).commit();
                            } else {
                                Helper.messageLogger(App.getAppContext(), Helper.LogType.INFO, "station", response.message());
                                if (response.code() == 401) {
//                                    getFragmentManager().beginTransaction().hide(tmp).commit();
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


    public void onButtonViewMapClick(View v) {
        MapsSelectFragment mapsSelectFragment = new MapsSelectFragment();
        String stationLocation = this.stationLocation.getText().toString();
        String coordinates[] = stationLocation.split("[, |]");
        if (coordinates.length == 2) {
            double posX = Double.parseDouble(coordinates[0]);
            double posY = Double.parseDouble(coordinates[1]);
            mapsSelectFragment.addMarker(posX, posY, "ok", true);
        } else {
            mapsSelectFragment.addMarker(51.10613247628298, 17.086756893213984, "", true);
        }
        getFragmentManager().beginTransaction().add(R.id.container, mapsSelectFragment).commit();
    }
}