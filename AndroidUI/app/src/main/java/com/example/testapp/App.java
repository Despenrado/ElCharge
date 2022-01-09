package com.example.testapp;

import android.app.Application;
import android.content.Context;
import android.location.LocationManager;
import android.widget.Toast;

import com.example.testapp.api.services.ElchargeService;
import com.example.testapp.greendao.DaoMaster;
import com.example.testapp.greendao.DaoSession;
import com.google.maps.GeoApiContext;

import org.greenrobot.greendao.database.Database;

public class App extends Application {

    private static Context context;
    private ElchargeService elchargeService; // retrofit + rx HTTP client for connection to backend
    private GeoApiContext geoApiContext;
    private DaoSession daoSession;

    @Override
    public void onCreate() {
        super.onCreate();
        App.context = getApplicationContext();
        DaoMaster.DevOpenHelper helper = new DaoMaster.DevOpenHelper(this, "elcharge-db");
        Database db = helper.getWritableDb();
        daoSession = new DaoMaster(db).newSession();
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

    public DaoSession getDaoSession() {
        return daoSession;
    }

    public Object getAppSystemService(String serviceName){
        return getSystemService(serviceName);
    }

}

