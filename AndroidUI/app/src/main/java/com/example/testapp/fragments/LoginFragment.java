package com.example.testapp.fragments;

import android.graphics.Color;
import android.os.Bundle;

import androidx.fragment.app.Fragment;

import android.util.Log;
import android.view.Gravity;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.Button;
import android.widget.EditText;
import android.widget.Toast;

import com.example.testapp.App;
import com.example.testapp.Helper;
import com.example.testapp.R;
import com.example.testapp.api.models.User;

import io.reactivex.android.schedulers.AndroidSchedulers;
import io.reactivex.disposables.CompositeDisposable;
import io.reactivex.observers.DisposableSingleObserver;
import io.reactivex.schedulers.Schedulers;
import retrofit2.Response;

public class LoginFragment extends Fragment {

    private CompositeDisposable disposable;
    private App app;

    @Override
    public void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        disposable = new CompositeDisposable();
        app = (App) getActivity().getApplication();
    }

    @Override
    public View onCreateView(LayoutInflater inflater, ViewGroup container,
                             Bundle savedInstanceState) {
        View view = inflater.inflate(R.layout.fragment_login, container, false);
        Button btnLogin = (Button) view.findViewById(R.id.buttonLogin);
        btnLogin.setOnClickListener(this::onButtonLoginClick);
        Button btnRegister = (Button) view.findViewById(R.id.buttonRegisterPage);
        btnRegister.setOnClickListener(this::onButtonRegisterPageClick);
        return view;
    }

    public void onButtonLoginClick(View v) {
        EditText email = (EditText) getView().findViewById(R.id.editTextEmail);
        EditText pass2 = (EditText) getView().findViewById(R.id.editTextPassword);
        if (email.getText().toString().equals("") || pass2.getText().toString().equals("")) {
            Helper.messageLogger(App.getAppContext(), Helper.LogType.NONE, "login", "Login or Password is incorrect");
        } else {
            User u = new User();
            u.setEmail(email.getText().toString());
            u.setPassword(pass2.getText().toString());
            login(u);
        }
    }


    public void onButtonRegisterPageClick(View v) {
        getFragmentManager().beginTransaction().add(R.id.container, new RegisterFragment()).commit();
        getFragmentManager().beginTransaction().remove(this).commit();
    }

    // send request to backend: login and get jwt token
    private void login(User user) {
        final LoginFragment tmpcls = this;
        disposable.add(app.getElchargeService().getUserApi().login(user)
                .subscribeOn(Schedulers.io())
                .observeOn(AndroidSchedulers.mainThread())
                .subscribeWith(new DisposableSingleObserver<Response<User>>() {
                    @Override
                    public void onSuccess(Response<User> response) {
                        try {
                            if (response.code() == 200) {
                                User tmp = response.body();
                                String token = response.headers().get("Authorization");
                                app.getElchargeService().setToken(token);
                                app.getElchargeService().setUser(tmp);
                                Helper.messageLogger(App.getAppContext(), Helper.LogType.INFO, "login", "LOGGED IN");
                                getFragmentManager().beginTransaction().remove(tmpcls).commit();
                            } else {
                                Helper.messageLogger(App.getAppContext(), Helper.LogType.INFO, "login", response.message());
                            }

                        } catch (Exception e) {
                            Helper.messageLogger(App.getAppContext(), Helper.LogType.ERR, "login", e.getMessage());
                        }
                    }

                    @Override
                    public void onError(Throwable e) {
                        Helper.messageLogger(App.getAppContext(), Helper.LogType.ERR, "login", e.getMessage());
                    }
                }));
    }

    @Override
    public void onDestroy() {
        disposable.dispose();
        super.onDestroy();
    }
}