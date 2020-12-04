package com.example.testapp.activities;

import androidx.annotation.NonNull;
import androidx.appcompat.app.AppCompatActivity;
import androidx.fragment.app.Fragment;

import android.os.Bundle;
import android.view.MenuItem;

import com.example.testapp.App;
import com.example.testapp.R;
import com.example.testapp.fragments.CreateStationFragment;
import com.example.testapp.fragments.MapsFragment;
import com.example.testapp.fragments.SearchStationFragment;
import com.example.testapp.fragments.SettingsFragment;
import com.google.android.material.bottomnavigation.BottomNavigationView;

public class MainActivity extends AppCompatActivity {

    private BottomNavigationView bottomNavigationView;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);
        bottomNavigationView = findViewById(R.id.bottomNav);
        bottomNavigationView.setOnNavigationItemSelectedListener(bottomNavMethod);
        getSupportFragmentManager().beginTransaction().replace(R.id.container, new MapsFragment()).commit();
    }

    private BottomNavigationView.OnNavigationItemSelectedListener bottomNavMethod = new
            BottomNavigationView.OnNavigationItemSelectedListener() {
                @Override
                public boolean onNavigationItemSelected(@NonNull MenuItem menuItem) {
                    Fragment fragment = null;
                    String tag = "";
                    switch (menuItem.getItemId()) {
                        case R.id.map:
                            tag = "main_maps";
                            fragment = getSupportFragmentManager().findFragmentByTag(tag);
                            if (fragment == null) {
                                fragment = new MapsFragment();
                            }
                            break;
                        case R.id.search:
                            tag = "search_station";
                            fragment = getSupportFragmentManager().findFragmentByTag(tag);
                            if (fragment == null) {
                                fragment = new SearchStationFragment();
                            }
                            break;
                        case R.id.create:
                            tag = "create_station";
                            fragment = getSupportFragmentManager().findFragmentByTag(tag);
                            if (fragment == null) {
                                fragment = new CreateStationFragment();
                            }
                            break;
                        case R.id.settings:
                            tag = "settings";
                            fragment = getSupportFragmentManager().findFragmentByTag(tag);
                            if (fragment == null) {
                                fragment = new SettingsFragment();
                            }
                            break;
                    }
                    getSupportFragmentManager().beginTransaction().replace(R.id.container, fragment, tag).commit();
                    return false;
                }
            };


}