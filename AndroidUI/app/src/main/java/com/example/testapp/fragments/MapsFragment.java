package com.example.testapp.fragments;

import androidx.annotation.NonNull;
import androidx.annotation.Nullable;
import androidx.annotation.RequiresPermission;
import androidx.core.app.ActivityCompat;
import androidx.core.content.ContextCompat;
import androidx.fragment.app.Fragment;

import android.Manifest;
import android.content.Context;
import android.content.pm.PackageManager;
import android.location.Location;
import android.location.LocationListener;
import android.location.LocationManager;
import android.os.Bundle;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;

import com.example.testapp.App;
import com.example.testapp.R;
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
import com.google.android.gms.tasks.OnSuccessListener;
import com.google.android.gms.tasks.Task;
import com.google.android.material.floatingactionbutton.FloatingActionButton;

import java.util.ArrayList;

public class MapsFragment extends Fragment implements OnMapReadyCallback {

    private GoogleMap googleMap;
    private FusedLocationProviderClient client;
    private ArrayList<MarkerOptions> markers;
    private Marker clientPositionMarker;
    private Marker selectedMarker;
    private App app;


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
        FloatingActionButton btnPlace = (FloatingActionButton) view.findViewById(R.id.floatingActionButtonPlace);
        btnPlace.setOnClickListener(this::onButtonPlaceClick);
        FloatingActionButton btnInfo = (FloatingActionButton) view.findViewById(R.id.floatingActionButtonInfo);
        btnInfo.setOnClickListener(this::onButtonInfoClick);
        FloatingActionButton btnMyLocation = (FloatingActionButton) view.findViewById(R.id.floatingActionButtonMyLocation);
        btnMyLocation.setOnClickListener(this::onButtonMyLocationClick);
    }


    @Override
    public void onMapReady(GoogleMap googleMap) {
        this.googleMap = setUpMarkers(googleMap);
        getPermissions();
    }

    private GoogleMap setUpMarkers(GoogleMap googleMap) {
        if (markers.size() != 0) {
            for (MarkerOptions marker : this.markers) {
                googleMap.addMarker(marker);
            }
            googleMap.moveCamera(CameraUpdateFactory.newLatLng(markers.get(0).getPosition()));
        } else {
            LatLng tmp = new LatLng(51.10613247628298, 17.086756893213984);
            googleMap.moveCamera(CameraUpdateFactory.newLatLng(tmp));
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
        LocationRequest locationRequest = new LocationRequest();
        locationRequest.setPriority(LocationRequest.PRIORITY_HIGH_ACCURACY);
        locationRequest.setFastestInterval(800);
        locationRequest.setInterval(1000);
        client.requestLocationUpdates(locationRequest, new LocationCallback() {
            @Override
            public void onLocationResult(LocationResult locationResult) {
                super.onLocationResult(locationResult);
                if (clientPositionMarker == null) {
                    clientPositionMarker = googleMap.addMarker(new MarkerOptions()
                            .title("you")
                            .position(new LatLng(locationResult.getLastLocation().getLatitude(), locationResult.getLastLocation().getLongitude()))
                            .icon(BitmapDescriptorFactory
                                    .defaultMarker(BitmapDescriptorFactory.HUE_BLUE)));
                } else {
                    clientPositionMarker.setPosition(new LatLng(locationResult.getLastLocation().getLatitude(), locationResult.getLastLocation().getLongitude()));
                }
            }
        }, app.getMainLooper());
    }

    private void onButtonPlaceClick(View v) {
        FloatingActionButton btnPlace = (FloatingActionButton) v.findViewById(R.id.floatingActionButtonPlace);
        btnPlace.hide();
        this.selectedMarker = googleMap.addMarker(new MarkerOptions()
                .position(googleMap.getCameraPosition().target)
                .draggable(true)
                .title("marker"));
    }

    private void onButtonInfoClick(View v) {
        LatLng latLng = selectedMarker.getPosition();
        getFragmentManager().beginTransaction().add(R.id.container, new InfoFragment(latLng.latitude, latLng.longitude), "station_info").commit();
    }

    private void onButtonMyLocationClick(View v) {
        if (googleMap != null && clientPositionMarker != null) {
            googleMap.moveCamera(CameraUpdateFactory.newLatLng(clientPositionMarker.getPosition()));
        }
    }

    public void addMarker(double x, double y, String title, boolean draggable) {
        this.markers.add(new MarkerOptions()
                .position(new LatLng(x, y))
                .draggable(draggable)
                .title(title));
    }

    private void getPermissions() {
        client = LocationServices.getFusedLocationProviderClient(App.getAppContext());
        if (ActivityCompat.checkSelfPermission(App.getAppContext(), Manifest.permission.ACCESS_FINE_LOCATION) != PackageManager.PERMISSION_GRANTED &&
                ContextCompat.checkSelfPermission(App.getAppContext(), Manifest.permission.ACCESS_COARSE_LOCATION) != PackageManager.PERMISSION_GRANTED) {
            ActivityCompat.requestPermissions(getActivity(), new String[]{Manifest.permission.ACCESS_COARSE_LOCATION, Manifest.permission.ACCESS_FINE_LOCATION}, 1);
            getPermissions();
        } else {
            getCurrentLocation();
        }
    }
}