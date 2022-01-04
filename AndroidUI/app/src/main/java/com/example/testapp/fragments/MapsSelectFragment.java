package com.example.testapp.fragments;

import android.os.Bundle;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;

import androidx.annotation.NonNull;
import androidx.annotation.Nullable;
import androidx.fragment.app.Fragment;

import com.example.testapp.R;
import com.google.android.gms.maps.CameraUpdateFactory;
import com.google.android.gms.maps.GoogleMap;
import com.google.android.gms.maps.OnMapReadyCallback;
import com.google.android.gms.maps.SupportMapFragment;
import com.google.android.gms.maps.model.LatLng;
import com.google.android.gms.maps.model.Marker;
import com.google.android.gms.maps.model.MarkerOptions;
import com.google.android.material.floatingactionbutton.FloatingActionButton;

import java.util.ArrayList;

public class MapsSelectFragment extends Fragment implements OnMapReadyCallback {

    private GoogleMap googleMap;
    private ArrayList<MarkerOptions> markers;
    private final float DEFAULT_ZOOM = 14;
    private final Double DEFAULT_LATITUDE = 51.10613247628298;
    private final Double DEFAULT_LONGITUDE = 17.086756893213984;

    public MapsSelectFragment() {
        this.markers = new ArrayList<MarkerOptions>();
    }

    @Nullable
    @Override
    public View onCreateView(@NonNull LayoutInflater inflater,
                             @Nullable ViewGroup container,
                             @Nullable Bundle savedInstanceState) {
        View view = inflater.inflate(R.layout.fragment_maps_create_station, container, false);

        return view;
    }

    @Override
    public void onViewCreated(@NonNull View view, @Nullable Bundle savedInstanceState) {
        super.onViewCreated(view, savedInstanceState);
        SupportMapFragment mapFragment =
                (SupportMapFragment) getChildFragmentManager().findFragmentById(R.id.mapView);
        if (mapFragment != null) {
            mapFragment.getMapAsync(this::onMapReady);
        }
        FloatingActionButton btnOK = (FloatingActionButton) view.findViewById(R.id.floatingActionButtonAddLocation);
        btnOK.setOnClickListener(this::onButtonOKClick);
    }


    @Override
    public void onMapReady(GoogleMap googleMap) {
        this.googleMap = setUpMarkers(googleMap);
    }

    private GoogleMap setUpMarkers(GoogleMap googleMap){
        if (markers.size() != 0) {
            for (MarkerOptions marker : this.markers) {
                googleMap.addMarker(marker);
                googleMap.setOnMarkerDragListener(new GoogleMap.OnMarkerDragListener() {
                    @Override
                    public void onMarkerDragStart(Marker marker) {
                    }

                    @Override
                    public void onMarkerDrag(Marker marker) {
                    }

                    @Override
                    public void onMarkerDragEnd(Marker marker) {
                        markers.get(0).position(marker.getPosition());
                    }
                });
            }
            googleMap.moveCamera(CameraUpdateFactory.newLatLngZoom(markers.get(0).getPosition(), DEFAULT_ZOOM));
        } else {
            googleMap.moveCamera(CameraUpdateFactory.newLatLngZoom(new LatLng(DEFAULT_LATITUDE, DEFAULT_LONGITUDE),DEFAULT_ZOOM));
        }
        return googleMap;
    }

    public void addMarker(double x, double y, String title, boolean draggable) {
        this.markers.add(new MarkerOptions()
                .position(new LatLng(x, y))
                .draggable(draggable)
                .title(title));
    }

    public void onButtonOKClick(View v) {
        CreateStationFragment createStationFragment = (CreateStationFragment)  getFragmentManager().findFragmentByTag("create_station");
        if(createStationFragment != null){
            createStationFragment.stationLocation.setText(Double.toString(markers.get(0).getPosition().latitude)+" "+Double.toString(markers.get(0).getPosition().longitude));
        }
        SearchStationFragment searchStationFragment = (SearchStationFragment)  getFragmentManager().findFragmentByTag("search_station");
        if(searchStationFragment != null){
            searchStationFragment.editTextSearch.setText(Double.toString(markers.get(0).getPosition().latitude)+" "+Double.toString(markers.get(0).getPosition().longitude));
        }
        getFragmentManager().beginTransaction().remove(this).commit();
    }
}