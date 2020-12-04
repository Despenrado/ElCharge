package com.example.testapp.fragments;

import android.Manifest;
import android.app.Activity;
import android.content.Context;
import android.content.pm.PackageManager;
import android.location.Location;
import android.location.LocationListener;
import android.location.LocationManager;
import android.os.Bundle;

import androidx.annotation.NonNull;
import androidx.core.app.ActivityCompat;
import androidx.core.content.ContextCompat;
import androidx.fragment.app.Fragment;

import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.AdapterView;
import android.widget.ArrayAdapter;
import android.widget.Button;
import android.widget.EditText;
import android.widget.SeekBar;
import android.widget.Spinner;
import android.widget.TextView;

import com.example.testapp.App;
import com.example.testapp.Helper;
import com.example.testapp.R;
import com.example.testapp.activities.MainActivity;
import com.example.testapp.api.models.Station;
import com.google.android.gms.location.FusedLocationProviderClient;
import com.google.android.gms.location.LocationServices;
import com.google.android.gms.tasks.OnSuccessListener;
import com.google.android.gms.tasks.Task;

import java.util.List;

import io.reactivex.android.schedulers.AndroidSchedulers;
import io.reactivex.disposables.CompositeDisposable;
import io.reactivex.observers.DisposableSingleObserver;
import io.reactivex.schedulers.Schedulers;
import retrofit2.Response;

public class SearchStationFragment extends Fragment implements AdapterView.OnItemSelectedListener {

    private ArrayAdapter arrayAdapter;

    private CompositeDisposable disposable = new CompositeDisposable();
    private App app;
    private Spinner spinnerSearchBy;
    private long searchBy = 0;
    private SeekBar seekBarRadius;
    private TextView textViewRadius;
    private int radius = 50;
    EditText editTextSearch;
    private Button buttonGo;
    private Activity activity;


    @Override
    public void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        app = (App) getActivity().getApplication();
    }

    @Override
    public View onCreateView(LayoutInflater inflater, ViewGroup container,
                             Bundle savedInstanceState) {
        View view = inflater.inflate(R.layout.fragment_search_station, container, false);
        arrayAdapter = ArrayAdapter.createFromResource(App.getAppContext(), R.array.search_by, android.R.layout.simple_spinner_item);
        arrayAdapter.setDropDownViewResource(android.R.layout.simple_spinner_dropdown_item);
        spinnerSearchBy = view.findViewById(R.id.spinnerSearch);
        spinnerSearchBy.setAdapter(arrayAdapter);
        spinnerSearchBy.setOnItemSelectedListener(this);
        textViewRadius = view.findViewById(R.id.textViewRadius);
        textViewRadius.setText("search radius: " + String.valueOf(radius) + " km");
        seekBarRadius = view.findViewById(R.id.seekBarDistance);
        seekBarRadius.setOnSeekBarChangeListener(new SeekBar.OnSeekBarChangeListener() {
            @Override
            public void onProgressChanged(SeekBar seekBar, int progress, boolean fromUser) {
                radius = seekBar.getProgress();
                textViewRadius.setText("search radius: " + String.valueOf(radius) + " km");
            }

            @Override
            public void onStartTrackingTouch(SeekBar seekBar) {
            }

            @Override
            public void onStopTrackingTouch(SeekBar seekBar) {
            }
        });
        editTextSearch = view.findViewById(R.id.editTextSearch);
        buttonGo = view.findViewById(R.id.buttonGo);
        buttonGo.setOnClickListener(this::onButtonGoClick);
        activity = getActivity();
        return view;
    }

    @Override
    public void onItemSelected(AdapterView<?> parent, View view, int position, long id) {
        searchBy = id;
        // init task location
//        Task<Location> locationTask = client.getLastLocation();
//        locationTask.addOnSuccessListener(new OnSuccessListener<Location>() {
//            @Override
//            public void onSuccess(Location location) {
//                if (location != null) {
//                    editTextSearch.setText(Double.toString(location.getLatitude()) + " " + Double.toString(location.getLongitude()));
//                }
//            }
//        });
        if (id == 2) {
            MapsSelectFragment mapsSelectFragment = new MapsSelectFragment();
            LocationManager locationManager = (LocationManager) app.getAppSystemService(Context.LOCATION_SERVICE);
            if (ContextCompat.checkSelfPermission(App.getAppContext(), Manifest.permission.ACCESS_COARSE_LOCATION) != PackageManager.PERMISSION_GRANTED &&
                    ContextCompat.checkSelfPermission(app.getApplicationContext(), Manifest.permission.ACCESS_FINE_LOCATION) != PackageManager.PERMISSION_GRANTED) {
                ActivityCompat.requestPermissions(activity, new String[]{Manifest.permission.ACCESS_COARSE_LOCATION, Manifest.permission.ACCESS_FINE_LOCATION}, 1);
            }
            if (ContextCompat.checkSelfPermission(App.getAppContext(), Manifest.permission.ACCESS_COARSE_LOCATION) == PackageManager.PERMISSION_GRANTED &&
                    ContextCompat.checkSelfPermission(app.getApplicationContext(), Manifest.permission.ACCESS_FINE_LOCATION) == PackageManager.PERMISSION_GRANTED) {
                FusedLocationProviderClient client = LocationServices.getFusedLocationProviderClient(App.getAppContext());
                Task<Location> locationTask = client.getLastLocation();
                locationTask.addOnSuccessListener(new OnSuccessListener<Location>() {
                    @Override
                    public void onSuccess(Location location) {
                        if (location != null) {
                            editTextSearch.setText(Double.toString(location.getLatitude()) + " " + Double.toString(location.getLongitude()));
                        }
                    }
                });
            }
        }
        if (id == 3) {
            MapsSelectFragment mapsSelectFragment = new MapsSelectFragment();
            String stationLocation = this.editTextSearch.getText().toString();
            String coordinates[] = stationLocation.split("[, |]");
            if (coordinates.length == 2) {
                double posX = Double.parseDouble(coordinates[0]);
                double posY = Double.parseDouble(coordinates[1]);
                mapsSelectFragment.addMarker(posX, posY, "ok", true);
            } else {
                mapsSelectFragment.addMarker(51.10613247628298, 17.086756893213984, "", true);
            }
            getFragmentManager().beginTransaction().add(R.id.container, mapsSelectFragment).commit();
            return;
        }
    }

    @Override
    public void onNothingSelected(AdapterView<?> parent) {
        searchBy = 0;
    }

    public void onButtonGoClick(View view) {
        if (searchBy == 0 || searchBy == 2 || searchBy == 3) {
            String stationLocation = this.editTextSearch.getText().toString();
            String coordinates[] = stationLocation.split("[, |]");
            if (coordinates.length < 2) {
                Helper.messageLogger(App.getAppContext(), Helper.LogType.NONE, "station", "wrong location");
                return;
            }
            try {
                double lat = Double.parseDouble(coordinates[0]);
                double lng = Double.parseDouble(coordinates[1]);
                findByLatAndLngAndDist(lat, lng, radius);
            } catch (NumberFormatException e) {
                Helper.messageLogger(App.getAppContext(), Helper.LogType.ERR, "station", e.getMessage());
            }
            return;
        }
        if (searchBy == 4) {
            String name = this.editTextSearch.getText().toString();
            findByName(name);
            return;
        }
        if (searchBy == 5) {
            String description = this.editTextSearch.getText().toString();
            findByDescription(description);
            return;
        }
    }

    private void findByLatAndLngAndDist(double lat, double lng, int dist) {
        disposable.add(app.getElchargeService().getStationApi().readStationsByLatAndLngAndDist(0, Integer.MAX_VALUE, lat, lng, dist)
                .subscribeOn(Schedulers.io())
                .observeOn(AndroidSchedulers.mainThread())
                .subscribeWith(new StationDisposableSingleObserver()));
    }

    private void findByName(String name) {
        disposable.add(app.getElchargeService().getStationApi().readStationsByName(0, Integer.MAX_VALUE, name)
                .subscribeOn(Schedulers.io())
                .observeOn(AndroidSchedulers.mainThread())
                .subscribeWith(new StationDisposableSingleObserver()));
    }

    private void findByDescription(String description) {
        disposable.add(app.getElchargeService().getStationApi().readStationsByDescription(0, Integer.MAX_VALUE, description)
                .subscribeOn(Schedulers.io())
                .observeOn(AndroidSchedulers.mainThread())
                .subscribeWith(new StationDisposableSingleObserver()));
    }

    private class StationDisposableSingleObserver extends DisposableSingleObserver<Response<List<Station>>> {
        @Override
        public void onSuccess(Response<List<Station>> response) {
            try {
                if (response.code() == 200) {
                    List<Station> stationList = response.body();
                    MapsFragment mapsFragment = new MapsFragment();
                    for (Station station : stationList) {
                        mapsFragment.addMarker(station.getLatitude(), station.getLongitude(), station.getStationName(), false);
                    }
                    Helper.messageLogger(App.getAppContext(), Helper.LogType.INFO, "station", response.message());
                    getFragmentManager().beginTransaction().replace(R.id.container, mapsFragment, "main_maps").commit();
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
    }
}
