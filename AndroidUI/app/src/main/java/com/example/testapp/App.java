package com.example.testapp;

import android.app.Application;
import android.content.Context;
import android.location.LocationManager;
import android.widget.Toast;

import com.example.testapp.api.services.ElchargeService;
import com.google.maps.GeoApiContext;

public class App extends Application {

    private static Context context;
    private ElchargeService elchargeService; // retrofit + rx HTTP client for connection to backend
    private GeoApiContext geoApiContext;

    @Override
    public void onCreate() {
        super.onCreate();
        App.context = getApplicationContext();
        elchargeService = new ElchargeService();
        geoApiContext = new GeoApiContext.Builder()
                .apiKey(getString(R.string.google_maps_key))
                .build();
    }

    public static Context getAppContext() {
        return context;
    }

    public ElchargeService getElchargeService() {
        try {
            return elchargeService;
        }catch (Exception e){
            Toast.makeText(context, "server not found", Toast.LENGTH_LONG);
            return null;
        }
    }

    public GeoApiContext getGeoApiContext() {
        if (geoApiContext == null){
            geoApiContext = new GeoApiContext.Builder()
                    .apiKey(getString(R.string.google_maps_key))
                    .build();
        }
        return geoApiContext;
    }

    public Object getAppSystemService(String serviceName){
        return getSystemService(serviceName);
    }

}

