package com.example.testapp.fragments;


import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.TextView;

import androidx.annotation.NonNull;
import androidx.recyclerview.widget.RecyclerView;

import com.example.testapp.Helper;
import com.example.testapp.R;
import com.example.testapp.api.models.Comment;

import java.text.DateFormat;
import java.text.SimpleDateFormat;
import java.util.ArrayList;
import java.util.Collections;
import java.util.Comparator;
import java.util.List;

public class CommentsAdapter extends RecyclerView.Adapter<CommentsAdapter.CommentViewHolder> {
    private List<Comment> commentList;

    public CommentsAdapter(List<Comment> commentList) {
//        Collections.sort(commentList, new Comparator<Comment>() {
//            public int compare(Comment o1, Comment o2) {
//                return o1.getUpdateAt().compareTo(o2.getUpdateAt());
//            }
//        });
        if (commentList != null) {
            this.commentList = commentList;
        }else {
            this.commentList = new ArrayList<>();
        }
    }

    @NonNull
    @Override
    public CommentsAdapter.CommentViewHolder onCreateViewHolder(@NonNull ViewGroup parent, int viewType) {
        View view = LayoutInflater.from(parent.getContext()).inflate(R.layout.info_view, parent, false);

        return new CommentViewHolder(view);
    }

    @Override
    public void onBindViewHolder(@NonNull CommentsAdapter.CommentViewHolder holder, int position) {
        holder.userName.setText(commentList.get(position).getUserName());
        holder.description.setText(commentList.get(position).getText());
        holder.rating.setText(String.format("%.2f", commentList.get(position).getRating()) + " / 5");
        holder.date.setText(Helper.getDateFromISO8601(commentList.get(position).getUpdateAt()));
    }

    @Override
    public int getItemCount() {
        return commentList.size();
    }

    public class CommentViewHolder extends RecyclerView.ViewHolder {
        TextView userName;
        TextView description;
        TextView rating;
        TextView date;

        public CommentViewHolder(@NonNull View itemView) {
            super(itemView);
            userName = itemView.findViewById(R.id.username);
            description = itemView.findViewById(R.id.description);
            rating = itemView.findViewById(R.id.rating);
            date = itemView.findViewById(R.id.date);
        }
    }
}
