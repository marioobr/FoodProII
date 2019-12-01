package ni.devotion.mvvm

import android.content.Intent
import androidx.appcompat.app.AppCompatActivity
import android.os.Bundle
import android.widget.Toast
import androidx.lifecycle.Observer
import androidx.recyclerview.widget.GridLayoutManager
import androidx.recyclerview.widget.RecyclerView
import kotlinx.android.synthetic.main.activity_main.*
import kotlinx.android.synthetic.main.activity_main1.*
import ni.abril.azb.megaboicotapp.ui.adapters.DepartmentAdapter
import ni.devotion.mvvm.adapters.RestaurantAdapter
import ni.devotion.mvvm.gui.MapsActivity
import ni.devotion.mvvm.gui.MenuActivity
import ni.devotion.mvvm.viewModel.DepartmentsViewModel
import ni.devotion.mvvm.viewModel.RestaurantViewModel
import org.koin.androidx.viewmodel.ext.android.viewModel

class MainActivity : AppCompatActivity() {


    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main1)

        loginBtn.setOnClickListener {
            val intent = Intent(this, MenuActivity::class.java)

            startActivity(intent)
        }

        mapIcon.setOnClickListener {
            val intent = Intent(this, MapsActivity::class.java)
            startActivity(intent)
        }
    }
}
