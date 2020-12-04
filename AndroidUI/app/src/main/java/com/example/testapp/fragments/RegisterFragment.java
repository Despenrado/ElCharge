package com.example.testapp.fragments;

import android.content.Intent;
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

public class RegisterFragment extends Fragment {

    private CompositeDisposable disposable = new CompositeDisposable();
    private App app;

    @Override
    public void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        app = (App) getActivity().getApplication();
    }

    @Override
    public View onCreateView(LayoutInflater inflater, ViewGroup container,
                             Bundle savedInstanceState) {
        View view = inflater.inflate(R.layout.fragment_register, container, false);
        Button btnRegister = (Button) view.findViewById(R.id.buttonRegister);
        btnRegister.setOnClickListener(this::onButtonRegisterClick);
        Button btnLogin = (Button) view.findViewById(R.id.buttonLoginPage);
        btnLogin.setOnClickListener(this::onButtonLoginPageClick);
        return view;
    }

    public void onButtonRegisterClick(View v) {
        EditText pass1 = (EditText) getView().findViewById(R.id.editTextPassword);
        EditText pass2 = (EditText) getView().findViewById(R.id.editTextPassword2);
        EditText email = (EditText) getView().findViewById(R.id.editTextEmail);
        EditText uname = (EditText) getView().findViewById(R.id.editTextUserName);
        if (pass1.getText().toString().equals(pass2.getText().toString()) && !email.getText().toString().equals("") && !uname.getText().toString().equals("")) {
                User u = new User();
                u.setEmail(email.getText().toString());
                u.setPassword(pass2.getText().toString());
                u.setUserName(uname.getText().toString());
                register(u);
        } else {
            Helper.messageLogger(App.getAppContext(), Helper.LogType.NONE, "register", "passwords are not equal");
        }
    }


    public void onButtonLoginPageClick(View v) {
        getFragmentManager().beginTransaction().add(R.id.container, new LoginFragment()).commit();
        getFragmentManager().beginTransaction().remove(this).commit();
    }

    // send request to backend: register and get jwt token
    private void register(User user) {
        final RegisterFragment tmpcls = this;
        disposable.add(app.getElchargeService().getUserApi().createUser(user)
                .subscribeOn(Schedulers.io())
                .observeOn(AndroidSchedulers.mainThread())
                .subscribeWith(new DisposableSingleObserver<Response<User>>() {
                    @Override
                    public void onSuccess(Response<User> response) {
                        try {
                            if (response.code() == 201) {
                                User tmp = response.body();
                                String token = response.headers().get("Authorization");
                                app.getElchargeService().setToken(token);
                                app.getElchargeService().setUser(tmp);
                                Helper.messageLogger(App.getAppContext(), Helper.LogType.INFO, "register", "account created successfully\n\nLOGGED IN");
                                getFragmentManager().beginTransaction().remove(tmpcls).commit();
                            }else{
                                Helper.messageLogger(App.getAppContext(), Helper.LogType.INFO, "register", response.message());
                            }

                        } catch (Exception e) {
                            Helper.messageLogger(App.getAppContext(), Helper.LogType.ERR, "register", e.getMessage());
                        }
                    }

                    @Override
                    public void onError(Throwable e) {
                        Helper.messageLogger(App.getAppContext(), Helper.LogType.ERR, "register", e.getMessage());
                    }
                }));
    }

    @Override
    public void onDestroy() {
        disposable.dispose();
        super.onDestroy();
    }
}