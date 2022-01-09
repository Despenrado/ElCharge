package com.example.testapp.fragments;

import android.Manifest;
import android.app.Activity;
import android.content.pm.PackageManager;
import android.location.Address;
import android.location.Geocoder;
import android.location.Location;
import android.os.Bundle;

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
import com.example.testapp.greendao.DaoSession;
import com.example.testapp.utils.Helper;
import com.example.testapp.R;
import com.example.testapp.models.Station;
import com.google.android.gms.location.FusedLocationProviderClient;
import com.google.android.gms.location.LocationServices;
import com.google.android.gms.tasks.OnSuccessListener;
import com.google.android.gms.tasks.Task;

import org.greenrobot.greendao.async.AsyncOperation;
import org.greenrobot.greendao.async.AsyncOperationListener;
import org.greenrobot.greendao.async.AsyncSession;

import java.io.IOException;
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
    private Geocoder geocoder;
    private final Double LATITUDE = 51.10613247628298;
    private final Double LONGITUDE = 17.086756893213984;
    private final float DEFAULT_ZOOM = 14;


    @Override
    public void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        app = (App) getActivity().getApplication();
        geocoder = new Geocoder(app);
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

    public String getUserLatLng() {
        if (ContextCompat.checkSelfPermission(App.getAppContext(), Manifest.permission.ACCESS_COARSE_LOCATION) != PackageManager.PERMISSION_GRANTED &&
                ContextCompat.checkSelfPermission(app.getApplicationContext(), Manifest.permission.ACCESS_FINE_LOCATION) != PackageManager.PERMISSION_GRANTED) {
            ActivityCompat.requestPermissions(activity, new String[]{Manifest.permission.ACCESS_COARSE_LOCATION, Manifest.permission.ACCESS_FINE_LOCATION}, 1);
        }
        if (ContextCompat.checkSelfPermission(App.getAppContext(), Manifest.permission.ACCESS_COARSE_LOCATION) == PackageManager.PERMISSION_GRANTED &&
                ContextCompat.checkSelfPermission(app.getApplicationContext(), Manifest.permission.ACCESS_FINE_LOCATION) == PackageManager.PERMISSION_GRANTED) {
            FusedLocationProviderClient client = LocationServices.getFusedLocationProviderClient(App.getAppContext());
            Task<Location> locationTask = client.getLastLocation();
            final String[] latLng = {""};
            locationTask.addOnSuccessListener(new OnSuccessListener<Location>() {
                @Override
                public void onSuccess(Location location) {
                    if (location != null) {
                        latLng[0] = Double.toString(location.getLatitude()) + " " + Double.toString(location.getLongitude());
                    }
                }
            });
            if (!latLng[0].equals("")) {
                return latLng[0];
            }
        }
        return LATITUDE.toString() + " " + LONGITUDE.toString();
    }

    @Override
    public void onItemSelected(AdapterView<?> parent, View view, int position, long id) {
        searchBy = id;
        this.editTextSearch.setEnabled(true);
        if (id == 2) {
            editTextSearch.setText(getUserLatLng());
            this.editTextSearch.setEnabled(false);
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
                mapsSelectFragment.addMarker(LATITUDE, LONGITUDE, "", true);
            }
            getFragmentManager().beginTransaction().add(R.id.container, mapsSelectFragment).commit();
            return;
        }
        if (id == 6) {
            this.editTextSearch.setText("");
            this.editTextSearch.setEnabled(false);
        }
    }

    @Override
    public void onNothingSelected(AdapterView<?> parent) {
        searchBy = 0;
    }

    public void onButtonGoClick(View view) {
        if (searchBy == 1) {
            try {
                Address location = geocoder.getFromLocationName(this.editTextSearch.getText().toString(), 1).get(0);
                findByLatAndLngAndDist(location.getLatitude(), location.getLongitude(), radius);
            } catch (IOException e) {
                Helper.messageLogger(App.getAppContext(), Helper.LogType.NONE, "search", e.getMessage());
            } catch (NullPointerException | IndexOutOfBoundsException e) {
                Helper.messageLogger(App.getAppContext(), Helper.LogType.NONE, "search", "Address NOT FOUND");
            }
        }
        if (searchBy == 0 || searchBy == 2 || searchBy == 3) {
            String stationLocation = this.editTextSearch.getText().toString();
            String coordinates[] = stationLocation.split("[, |]");
            if (coordinates.length < 2) {
                Helper.messageLogger(App.getAppContext(), Helper.LogType.NONE, "search", "wrong location");
                return;
            }
            try {
                double lat = Double.parseDouble(coordinates[0]);
                double lng = Double.parseDouble(coordinates[1]);
                findByLatAndLngAndDist(lat, lng, radius);
            } catch (NumberFormatException e) {
                Helper.messageLogger(App.getAppContext(), Helper.LogType.ERR, "search", e.getMessage());
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
        if (searchBy == 6) {
            getAllStationsFromSQLite();
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

    private void getAllStationsFromSQLite() {
        AsyncSession asyncSession = app.getDaoSession().startAsyncSession();
        asyncSession.setListener(new AsyncOperationListener() {
            @Override
            public void onAsyncOperationCompleted(AsyncOperation operation) {
                List<Station> stationList = (List<Station>) operation.getResult();
                MapsFragment mapsFragment = new MapsFragment();
                for (Station station : stationList) {
                    mapsFragment.addMarker(station.getLatitude(), station.getLongitude(), station.getStationName(), false);
                }
                getFragmentManager().beginTransaction().replace(R.id.container, mapsFragment, "main_maps").commit();
            }
        });
        asyncSession.loadAll(Station.class);
    }

    private class StationDisposableSingleObserver extends DisposableSingleObserver<Response<List<Station>>> {
        @Override
        public void onSuccess(Response<List<Station>> response) {
            try {
                if (response.code() == 200) {
                    List<Station> stationList = response.body();
                    MapsFragment mapsFragment = new MapsFragment();
                    DaoSession daoSession = app.getDaoSession();
                    for (Station station : stationList) {
                        daoSession.insertOrReplace(station);
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

    @Override
    public void onDestroy() {
        disposable.dispose();
        super.onDestroy();
    }
}
