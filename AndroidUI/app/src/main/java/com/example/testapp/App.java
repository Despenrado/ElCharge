package com.example.testapp;

import android.app.Application;
import android.content.Context;
import android.location.LocationManager;
import android.widget.Toast;

import com.example.testapp.api.services.ElchargeService;

public class App extends Application {

    private static Context context;
    private ElchargeService elchargeService; // retrofit + rx HTTP client for connection to backend

    @Override
    public void onCreate() {
        super.onCreate();
        App.context = getApplicationContext();
        elchargeService = new ElchargeService();
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

    public Object getAppSystemService(String serviceName){
        return getSystemService(serviceName);
    }

}

