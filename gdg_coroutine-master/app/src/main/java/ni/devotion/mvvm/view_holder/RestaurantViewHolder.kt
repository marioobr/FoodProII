package ni.devotion.mvvm.view_holder

import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import androidx.recyclerview.widget.RecyclerView
import kotlinx.android.extensions.LayoutContainer
import kotlinx.android.synthetic.main.restaurant_content.*
import ni.devotion.mvvm.R
import ni.devotion.mvvm.model.Restaurant

class RestaurantViewHolder constructor
    (override val containerView: View) : RecyclerView.ViewHolder(containerView),
    LayoutContainer {

    fun bind(rest: Restaurant) {
        restName.text = rest.name
        restAddress.text = rest.address
        restPhone.text = rest.phone.toString()
    }

    companion object {
        fun create(parent: ViewGroup): RestaurantViewHolder {
            return RestaurantViewHolder(LayoutInflater.from
                (parent.context).inflate(
                R.layout.restaurant_card, parent, false)
            )
        }
    }
}