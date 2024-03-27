package models
import models.Products
import scala.collection.mutable.ListBuffer

case class Category(var name: String, var products: ListBuffer[Products], var index: Int)
