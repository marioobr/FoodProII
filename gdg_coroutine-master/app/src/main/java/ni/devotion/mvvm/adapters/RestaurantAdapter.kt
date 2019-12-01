package ni.devotion.mvvm.adapters

import android.view.ViewGroup
import androidx.recyclerview.widget.DiffUtil
import androidx.recyclerview.widget.ListAdapter
import ni.devotion.mvvm.model.Restaurant
import ni.devotion.mvvm.view_holder.RestaurantViewHolder

class RestaurantAdapter : ListAdapter<Restaurant, RestaurantViewHolder>(DIFF_CALLBACK){
    companion object{
        private val DIFF_CALLBACK = object : DiffUtil.ItemCallback<Restaurant>() {
            override fun areItemsTheSame(oldItem: Restaurant, newItem: Restaurant) = oldItem.id == newItem.id
            override fun areContentsTheSame(oldItem: Restaurant, newItem: Restaurant) = oldItem == newItem
        }
    }

    override fun onCreateViewHolder(
        parent: ViewGroup,
        viewType: Int) =
        RestaurantViewHolder.create(
            parent
        )

    override fun onBindViewHolder(holder: RestaurantViewHolder, position: Int) {
        holder.bind(getItem(position))
    }
}