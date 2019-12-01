package ni.devotion.mvvm.model

import androidx.annotation.Keep

@Keep
data class Combo(
    val id: Int,
    val name: String,
    val description: String,
    val price: Int,
    val image: String
)