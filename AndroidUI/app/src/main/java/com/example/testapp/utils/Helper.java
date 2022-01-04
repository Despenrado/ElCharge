package com.example.testapp.utils;

import android.content.Context;
import android.content.SharedPreferences;
import android.content.res.Resources;
import android.util.Log;
import android.view.Gravity;
import android.widget.Toast;

import com.example.testapp.R;

import java.io.FileOutputStream;
import java.io.IOException;
import java.io.InputStream;
import java.text.ParseException;
import java.text.SimpleDateFormat;
import java.util.Date;
import java.util.Properties;

public final class Helper {
    private static final String TAG = "Helper";

    public enum LogType {
        ERR,
        INFO,
        NONE,
    }


    public static String getConfigValue(Context context, String name) {
        try {
            Resources resources = context.getResources();
            InputStream rawResource = resources.openRawResource(R.raw.config);
            Properties properties = new Properties();
            properties.load(rawResource);
            rawResource.close();
            return properties.getProperty(name);
        } catch (Resources.NotFoundException e) {
            Log.e(TAG, "Unable to find the config file: " + e.getMessage());
        } catch (IOException e) {
            Log.e(TAG, "Failed to open config file\n" + e.getMessage());
        }
        return null;
    }

    public static void saveToSharedPreferences(Context context, String name, String value) {
        SharedPreferences sharedPref = context.getSharedPreferences("com.example.testapp", Context.MODE_PRIVATE);
        SharedPreferences.Editor editor = sharedPref.edit();
        editor.putString(name, value);
        editor.apply();
    }

    public static String readFromSharedPreferences(Context context, String name) {
        SharedPreferences sharedPref = context.getSharedPreferences("com.example.testapp", Context.MODE_PRIVATE);
        return sharedPref.getString(name, "");
    }

    public static void messageLogger(Context context, LogType logType,String tag, String msg) {
        switch (logType){
            case ERR:
                Log.e(tag, msg);
                break;
            case INFO:
                Log.i(tag, msg);
                break;
            case NONE:
                break;
        }
        Toast toast = Toast.makeText(context, msg, Toast.LENGTH_LONG);
        toast.setGravity(Gravity.TOP, 0, 0);
        toast.show();
    }

    public static String getDateFromISO8601(String date){
        String outputDate;
        try {
            SimpleDateFormat simpledateformat = new SimpleDateFormat("yyyy-MM-dd'T'HH:mm:ssXXX");
            Date tempDate = simpledateformat.parse(date);
            outputDate = tempDate.toString();
        }catch (ParseException ex)
        {
            outputDate = "NoN";
        }
        return outputDate;
    }

}