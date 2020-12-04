package com.example.testapp.fragments;

import android.os.Bundle;

import androidx.annotation.NonNull;
import androidx.annotation.Nullable;
import androidx.fragment.app.Fragment;

import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.Button;
import android.widget.EditText;
import android.widget.RatingBar;

import com.example.testapp.App;
import com.example.testapp.Helper;
import com.example.testapp.R;
import com.example.testapp.api.models.Comment;
import com.example.testapp.api.models.User;
import com.google.android.material.floatingactionbutton.FloatingActionButton;

import io.reactivex.android.schedulers.AndroidSchedulers;
import io.reactivex.disposables.CompositeDisposable;
import io.reactivex.observers.DisposableSingleObserver;
import io.reactivex.schedulers.Schedulers;
import retrofit2.Response;


public class CreateCommentFragment extends Fragment {

    private CompositeDisposable disposable = new CompositeDisposable();
    private App app;
    private String stationID;

    public CreateCommentFragment(String stationID) {
        this.stationID = stationID;
    }

    @Override
    public void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
    }

    @Override
    public View onCreateView(LayoutInflater inflater, ViewGroup container,
                             Bundle savedInstanceState) {
        return inflater.inflate(R.layout.fragment_create_comment, container, false);
    }

    @Override
    public void onViewCreated(@NonNull View view, @Nullable Bundle savedInstanceState) {
        super.onViewCreated(view, savedInstanceState);
        this.app = (App) getActivity().getApplication();
        Button btnSave = (Button)getView().findViewById(R.id.buttonSave);
        btnSave.setOnClickListener(this::onButtonSave);
        Button btnCancel = (Button)getView().findViewById(R.id.buttonCancel);
        btnCancel.setOnClickListener(this::onButtonCancel);
    }

    private void onButtonSave(View v){
        Comment comment = new Comment();
        comment.setRating(((RatingBar)getView().findViewById(R.id.ratingBar)).getRating());
        comment.setText(((EditText)getView().findViewById(R.id.editTextText)).getText().toString());
        comment.setUserID(app.getElchargeService().getUser().getId());
        comment.setUserName(app.getElchargeService().getUser().getUserName());
        createComment(comment);
    }

    private void onButtonCancel(View v){
        getFragmentManager().beginTransaction().remove(this).commit();
    }

    private void createComment(Comment comment){
        final CreateCommentFragment tmpcls = this;
        disposable.add(app.getElchargeService().getCommentApi().createComment(stationID,comment)
                .subscribeOn(Schedulers.io())
                .observeOn(AndroidSchedulers.mainThread())
                .subscribeWith(new DisposableSingleObserver<Response<Comment>>() {
                    @Override
                    public void onSuccess(Response<Comment> response) {
                        try {
                            if (response.code() == 201) {
                                Helper.messageLogger(App.getAppContext(), Helper.LogType.INFO, "comment", response.message());
                                InfoFragment inf = (InfoFragment)getFragmentManager().findFragmentByTag("station_info");
                                if(inf != null){
                                    inf.updateInfo();
                                }
                                getFragmentManager().beginTransaction().remove(tmpcls).commit();
                            } else {
                                Helper.messageLogger(App.getAppContext(), Helper.LogType.INFO, "comment", response.message());
                                if (response.code() == 401) {
//                                    getFragmentManager().beginTransaction().hide(tmp).commit();
                                    LoginFragment lf = new LoginFragment();
                                    getFragmentManager().beginTransaction().add(R.id.container, lf).commit();
                                    getFragmentManager().beginTransaction().show(lf).commit();
                                }
                            }

                        } catch (Exception e) {
                            Helper.messageLogger(App.getAppContext(), Helper.LogType.ERR, "comment", e.getMessage());
                        }
                    }

                    @Override
                    public void onError(Throwable e) {
                        Helper.messageLogger(App.getAppContext(), Helper.LogType.ERR, "comment", e.getMessage());
                    }
                }));
    }

}