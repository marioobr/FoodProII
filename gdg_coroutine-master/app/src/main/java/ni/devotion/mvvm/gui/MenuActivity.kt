package ni.devotion.mvvm.gui

import androidx.appcompat.app.AppCompatActivity
import android.os.Bundle
import android.widget.Toast
import androidx.lifecycle.Observer
import androidx.recyclerview.widget.GridLayoutManager
import kotlinx.android.synthetic.main.activity_main.*
import ni.devotion.mvvm.R
import ni.devotion.mvvm.adapters.RestaurantAdapter
import ni.devotion.mvvm.viewModel.RestaurantViewModel
import org.koin.androidx.viewmodel.ext.android.viewModel

class MenuActivity : AppCompatActivity() {

    private val restsViewModel: RestaurantViewModel by viewModel()
    private val adapter: RestaurantAdapter by lazy { RestaurantAdapter() }

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)
        rvRestaurants.apply {
            layoutManager = GridLayoutManager(context, 2)
            setHasFixedSize(true)
            adapter = this@MenuActivity.adapter
        }
        restsViewModel.uiState.observe(this, Observer {
            val dataState = it ?: return@Observer
            if (!dataState.showProgress){ }
            if (dataState.restaurant != null && !dataState.restaurant.consumed){
                dataState.restaurant.consume()?.let { restaurant ->
                    adapter.submitList(restaurant)
                }
            }
            if (dataState.error != null && !dataState.error.consumed){
                dataState.error.consume()?.let { error ->
                    Toast.makeText(applicationContext, resources.getString(error), Toast.LENGTH_LONG).show()
                }
            }
        })
    }
}
