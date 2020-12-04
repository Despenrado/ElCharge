package com.example.testapp.fragments;

import android.os.Bundle;

import androidx.annotation.NonNull;
import androidx.fragment.app.Fragment;
import androidx.fragment.app.ListFragment;

import android.util.Log;
import android.view.Gravity;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.ArrayAdapter;
import android.widget.ListView;
import android.widget.Toast;

import com.example.testapp.App;
import com.example.testapp.Helper;
import com.example.testapp.R;
import com.example.testapp.api.models.User;

import io.reactivex.android.schedulers.AndroidSchedulers;
import io.reactivex.disposables.CompositeDisposable;
import io.reactivex.observers.DisposableSingleObserver;
import io.reactivex.schedulers.Schedulers;
import okhttp3.ResponseBody;

public class SettingsFragment extends ListFragment {

    private CompositeDisposable disposable = new CompositeDisposable();
    private App app;

    public SettingsFragment() {
        // Required empty public constructor
    }

    private String[] data = new String[]{
            "sign in",
            "log out"
    };

    @Override
    public void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        ArrayAdapter<String> adapter;
        adapter = new ArrayAdapter<String>(App.getAppContext(),
                android.R.layout.simple_list_item_1,
                data);
        setListAdapter(adapter);
        app = (App) getActivity().getApplication();
    }

    @Override
    public View onCreateView(LayoutInflater inflater, ViewGroup container,
                             Bundle savedInstanceState) {
        return inflater.inflate(R.layout.fragment_settings, container, false);
    }

    @Override
    public void onListItemClick(@NonNull ListView l, @NonNull View v, int position, long id) {
        switch ((String) getListAdapter().getItem(position)) {
            case "sign in":
                getFragmentManager().beginTransaction().add(R.id.container, new LoginFragment()).commit();
                break;
            case "log out":
                logout();
                break;
        }
    }

    // send request to backend: logout
    private void logout(){
        disposable.add(app.getElchargeService().getUserApi().logout(app.getElchargeService().getUser().getId())
                .subscribeOn(Schedulers.io())
                .observeOn(AndroidSchedulers.mainThread())
                .subscribeWith(new DisposableSingleObserver<ResponseBody>() {
                    @Override
                    public void onSuccess(ResponseBody responseBody) {
                        try {
                            if (responseBody != null) {
                                app.getElchargeService().setToken(""); //remove token
                                getFragmentManager().beginTransaction().replace(R.id.container, new LoginFragment()).commit();
                            }
                            Helper.messageLogger(App.getAppContext(), Helper.LogType.INFO, "logout", responseBody.string());
                        } catch (Exception e) {
                            Helper.messageLogger(App.getAppContext(), Helper.LogType.ERR, "logout", e.getMessage());
                        }
                    }

                    @Override
                    public void onError(Throwable e) {
                        Helper.messageLogger(App.getAppContext(), Helper.LogType.ERR, "logout", e.getMessage());
                    }
                }));
    }
}