package com.example.testapp.fragments;

import androidx.annotation.NonNull;
import androidx.annotation.Nullable;
import androidx.annotation.RequiresPermission;
import androidx.core.app.ActivityCompat;
import androidx.core.content.ContextCompat;
import androidx.fragment.app.Fragment;

import android.Manifest;
import android.content.pm.PackageManager;
import android.location.Location;
import android.os.Bundle;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;

import com.example.testapp.App;
import com.example.testapp.R;
import com.example.testapp.utils.Helper;
import com.google.android.gms.location.FusedLocationProviderClient;
import com.google.android.gms.location.LocationCallback;
import com.google.android.gms.location.LocationRequest;
import com.google.android.gms.location.LocationResult;
import com.google.android.gms.location.LocationServices;
import com.google.android.gms.maps.CameraUpdateFactory;
import com.google.android.gms.maps.GoogleMap;
import com.google.android.gms.maps.OnMapReadyCallback;
import com.google.android.gms.maps.SupportMapFragment;
import com.google.android.gms.maps.model.BitmapDescriptorFactory;
import com.google.android.gms.maps.model.LatLng;
import com.google.android.gms.maps.model.Marker;
import com.google.android.gms.maps.model.MarkerOptions;
import com.google.android.gms.maps.model.Polygon;
import com.google.android.gms.maps.model.Polyline;
import com.google.android.gms.maps.model.PolylineOptions;
import com.google.android.gms.tasks.OnCompleteListener;
import com.google.android.gms.tasks.Task;
import com.google.android.material.floatingactionbutton.FloatingActionButton;
import com.google.maps.DirectionsApiRequest;
import com.google.maps.PendingResult;
import com.google.maps.internal.PolylineEncoding;
import com.google.maps.model.DirectionsResult;
import com.google.maps.model.DirectionsRoute;


import java.util.ArrayList;
import java.util.List;
import java.util.Observable;

import io.reactivex.Completable;
import io.reactivex.CompletableObserver;
import io.reactivex.Observer;
import io.reactivex.Single;
import io.reactivex.android.schedulers.AndroidSchedulers;
import io.reactivex.disposables.CompositeDisposable;
import io.reactivex.disposables.Disposable;
import io.reactivex.functions.Action;
import io.reactivex.observers.DisposableCompletableObserver;
import io.reactivex.schedulers.Schedulers;

public class MapsFragment extends Fragment implements OnMapReadyCallback{

    private CompositeDisposable disposable = new CompositeDisposable();
    private GoogleMap googleMap;
    private FusedLocationProviderClient client;
    private ArrayList<MarkerOptions> markers;
    private Marker clientPositionMarker;
    private Marker selectedMarker;
    private App app;
    private final float DEFAULT_ZOOM = 14;
    private final Double DEFAULT_LATITUDE = 51.10613247628298;
    private final Double DEFAULT_LONGITUDE = 17.086756893213984;


    public MapsFragment() {
        this.markers = new ArrayList<MarkerOptions>();
    }

    @Nullable
    @Override
    public View onCreateView(@NonNull LayoutInflater inflater,
                             @Nullable ViewGroup container,
                             @Nullable Bundle savedInstanceState) {
        return inflater.inflate(R.layout.fragment_maps, container, false);
    }

    @Override
    public void onViewCreated(@NonNull View view, @Nullable Bundle savedInstanceState) {
        super.onViewCreated(view, savedInstanceState);
        SupportMapFragment mapFragment = (SupportMapFragment) getChildFragmentManager().findFragmentById(R.id.mapView);
        app = (App) getActivity().getApplication();
        if (mapFragment != null) {
            mapFragment.getMapAsync(this::onMapReady);
        }
//        FloatingActionButton btnPlace = (FloatingActionButton) view.findViewById(R.id.floatingActionButtonPlace);
//        btnPlace.setOnClickListener(this::onButtonPlaceClick);
        FloatingActionButton btnInfo = (FloatingActionButton) view.findViewById(R.id.floatingActionButtonInfo);
        btnInfo.setOnClickListener(this::onButtonInfoClick);
        FloatingActionButton btnMyLocation = (FloatingActionButton) view.findViewById(R.id.floatingActionButtonMyLocation);
        btnMyLocation.setOnClickListener(this::onButtonMyLocationClick);
        FloatingActionButton btnNavigate = (FloatingActionButton) view.findViewById(R.id.floatingActionNavigate);
        btnNavigate.setOnClickListener(this::onButtonDrawPolylineClick);
    }


    @Override
    public void onMapReady(GoogleMap googleMap) {
        this.googleMap = setUpMarkers(googleMap);
//        geoApiContext = new GeoApiContext.Builder()
//                .apiKey(getString(R.string.google_maps_key))
//                .build();
        getPermissionsAndShowUserPosition();
    }

    private GoogleMap setUpMarkers(GoogleMap googleMap) {
        if (markers.size() != 0) {
            for (MarkerOptions marker : this.markers) {
                googleMap.addMarker(marker);
            }
            googleMap.moveCamera(CameraUpdateFactory.newLatLngZoom(markers.get(0).getPosition(), DEFAULT_ZOOM));
        } else {
            googleMap.moveCamera(CameraUpdateFactory.newLatLngZoom(new LatLng(DEFAULT_LATITUDE, DEFAULT_LONGITUDE), DEFAULT_ZOOM));
        }
        googleMap.setOnMarkerClickListener(new GoogleMap.OnMarkerClickListener() {
            @Override
            public boolean onMarkerClick(Marker marker) {
                selectedMarker = marker;
                return false;
            }
        });
        return googleMap;
    }

    @RequiresPermission(
            anyOf = {"android.permission.ACCESS_COARSE_LOCATION", "android.permission.ACCESS_FINE_LOCATION"}
    )
    private void getCurrentLocation() {
        googleMap.setMyLocationEnabled(true);
//        client.getCurrentLocation()
//        client.getLastLocation().addOnCompleteListener(new OnCompleteListener<Location>() {
//            @Override
//            public void onComplete(@NonNull Task<Location> task) {
//                if (task.isSuccessful()) {
//                    Location location = task.getResult();
//                    clientPositionMarker = new MarkerOptions()
//                            .title("you")
//                            .position(new LatLng(location.getLatitude(), location.getLongitude()))
//                            .icon(BitmapDescriptorFactory
//                                    .defaultMarker(BitmapDescriptorFactory.HUE_BLUE));
//                }
//
//            }
//        });
        LocationRequest locationRequest = new LocationRequest();
        locationRequest.setPriority(LocationRequest.PRIORITY_HIGH_ACCURACY);
        locationRequest.setFastestInterval(100);
        locationRequest.setInterval(500);
        client.requestLocationUpdates(locationRequest, new LocationCallback() {
            @Override
            public void onLocationResult(LocationResult locationResult) {
                super.onLocationResult(locationResult);
                if (clientPositionMarker == null) {
                    clientPositionMarker = googleMap.addMarker(new MarkerOptions()
                            .title("you")
                            .position(new LatLng(locationResult.getLastLocation().getLatitude(), locationResult.getLastLocation().getLongitude()))
                            .visible(false));
//                    Helper.messageLogger(getContext(), Helper.LogType.INFO, "LatLng", locationResult.getLastLocation().getLatitude() + " " + locationResult.getLastLocation().getLongitude());
                } else {
                    clientPositionMarker.setPosition(new LatLng(locationResult.getLastLocation().getLatitude(), locationResult.getLastLocation().getLongitude()));
//                    Helper.messageLogger(getContext(), Helper.LogType.INFO, "LatLng", locationResult.getLastLocation().getLatitude() + " " + locationResult.getLastLocation().getLongitude());
                }
            }
        }, app.getMainLooper());
    }

//    private void onButtonPlaceClick(View v) {
//        FloatingActionButton btnPlace = (FloatingActionButton) v.findViewById(R.id.floatingActionButtonPlace);
//        btnPlace.hide();
//        this.selectedMarker = googleMap.addMarker(new MarkerOptions()
//                .position(googleMap.getCameraPosition().target)
//                .draggable(true)
//                .title("marker"));
//    }

    private void onButtonDrawPolylineClick(View v) {
        if (selectedMarker != null) {
            getPermissionsAndShowUserPosition();
            if (clientPositionMarker == null) {
                try {
                    Thread.sleep(1000);
                } catch (InterruptedException e) {
                }
            }
            if (clientPositionMarker != null && selectedMarker != null) {
                DirectionsApiRequest directionsApiRequest = new DirectionsApiRequest(app.getGeoApiContext());
                directionsApiRequest.alternatives(true);
                directionsApiRequest.origin(new com.google.maps.model.LatLng(clientPositionMarker.getPosition().latitude, clientPositionMarker.getPosition().longitude));
                directionsApiRequest.destination(new com.google.maps.model.LatLng(selectedMarker.getPosition().latitude, selectedMarker.getPosition().longitude));
                directionsApiRequest.setCallback(new PendingResult.Callback<DirectionsResult>() {
                    @Override
                    public void onResult(DirectionsResult result) {
                        disposable.add(Completable.fromAction(new Action() {
                            @Override
                            public void run() throws Exception {
                                for (DirectionsRoute route : result.routes) {
                                    List<LatLng> newDecodedPath = new ArrayList<>();
                                    List<com.google.maps.model.LatLng> decodedPath = PolylineEncoding.decode(route.overviewPolyline.getEncodedPath());
                                    for (com.google.maps.model.LatLng latLng : decodedPath) {
                                        newDecodedPath.add(new LatLng(latLng.lat, latLng.lng));
                                    }
                                    System.out.println(newDecodedPath.toString());
                                    getActivity().runOnUiThread(new Runnable() {
                                        @Override
                                        public void run() {
                                            Polyline polyline = googleMap.addPolyline(new PolylineOptions()
                                                    .clickable(true)
                                                    .addAll(newDecodedPath));
                                            polyline.setColor(ContextCompat.getColor(getActivity(), R.color.lightGoogleBlue));
//                                            googleMap.setOnPolylineClickListener(MapsFragment.this);
//                                            googleMap.setOnPolygonClickListener(MapsFragment.this);
                                        }
                                    });
                                }
                            }
                        }).subscribeOn(Schedulers.io())
                                .observeOn(AndroidSchedulers.mainThread())
                                .subscribe());
//                                .subscribeWith(new DisposableCompletableObserver() {
//
//                                    @Override
//                                    public void onComplete() {
////                        googleMap.addPolyline();
//                                    }
//
//                                    @Override
//                                    public void onError(@NonNull Throwable e) {
//
//                                    }
//                                })
//                        );
                    }

                    @Override
                    public void onFailure(Throwable e) {
                        Helper.messageLogger(App.getAppContext(), Helper.LogType.ERR, "directions", e.getMessage());
                    }
                });
            }
        }
    }

    private void onButtonInfoClick(View v) {
        if (selectedMarker != null) {
            LatLng latLng = selectedMarker.getPosition();
            getFragmentManager().beginTransaction().add(R.id.container, new InfoFragment(latLng.latitude, latLng.longitude), "station_info").commit();
        }
    }

    private void onButtonMyLocationClick(View v) {
        if (googleMap != null && clientPositionMarker != null) {
            googleMap.moveCamera(CameraUpdateFactory.newLatLngZoom(clientPositionMarker.getPosition(), DEFAULT_ZOOM));
        }
    }

    public void addMarker(double x, double y, String title, boolean draggable) {
        this.markers.add(new MarkerOptions()
                .position(new LatLng(x, y))
                .draggable(draggable)
                .title(title));
    }

    private void getPermissionsAndShowUserPosition() {
        client = LocationServices.getFusedLocationProviderClient(App.getAppContext());
        if (ActivityCompat.checkSelfPermission(App.getAppContext(), Manifest.permission.ACCESS_FINE_LOCATION) != PackageManager.PERMISSION_GRANTED &&
                ContextCompat.checkSelfPermission(App.getAppContext(), Manifest.permission.ACCESS_COARSE_LOCATION) != PackageManager.PERMISSION_GRANTED) {
            ActivityCompat.requestPermissions(getActivity(), new String[]{Manifest.permission.ACCESS_COARSE_LOCATION, Manifest.permission.ACCESS_FINE_LOCATION}, 1);
            try {
                Thread.sleep(100);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
            getPermissionsAndShowUserPosition();
        } else {
            getCurrentLocation();

        }
    }

//    @Override
//    public void onDestroy() {
//        disposable.dispose();
//        super.onDestroy();
//    }
}