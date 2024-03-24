package controllers

import models.Products
import javax.inject._
import play.api._
import play.api.mvc._
import scala.collection.mutable.ListBuffer
import scala.util.Try

/**
 * This controller creates an `Action` to handle HTTP requests to the
 * application's home page.
 */
@Singleton
class HomeController @Inject()(val controllerComponents: ControllerComponents) extends BaseController {

  /**
   * Create an Action to render an HTML page.
   *
   * The configuration in the `routes` file means that this method
   * will be called when the application receives a `GET` request with
   * a path of `/`.
   */
  def index() = Action { implicit request: Request[AnyContent] =>
    Ok(views.html.index())
  }

  val prodlist = new ListBuffer[Products]()
  prodlist += new Products(pname="VeryDumbName", pprice=12, pman="Debvils")

  def listsize(idcheck: Int): Boolean = {
    var sizer = prodlist.length
    var ok = false
    if (sizer > idcheck && idcheck >= 0) {
      ok = true
    }
    return ok
  }
  
  def pgetall() = Action { implicit request =>
    Ok(views.html.pgetall(prodlist))
  }

  def pgetone() = Action { implicit request =>
    val index = request.getQueryString("pid").flatMap(s => Try(s.toInt).toOption)
    var ok=false
    if (index.isDefined) {
      ok = listsize(index.get)
    }
    if (ok == true) {
      val product = index.flatMap(i => prodlist.lift(i))
      Ok(views.html.pgetone(product))
    }
    else

      BadGateway(s"Nie udało się oðebrać produktu o podanym indeksie")
  }

  def premoval() = Action { implicit request =>
    val index = request.getQueryString("prid").flatMap(s => Try(s.toInt).toOption)
    var ok=false
    if (index.isDefined) {
      ok = listsize(index.get)
    }
    if (ok == true) {
      val nice=index.get
      prodlist.remove(nice)
      //prodlist.patch(index.get, Nil, 1)
      Ok(s"Element $nice usunięty")
    }
    else
      BadGateway(s"Nie udało się usunąć elementu o podanym indeksie")
  }
  def pcreate() = Action { implicit request =>
    val name = request.getQueryString("pname").flatMap(s => Try(s.toString).toOption)
    val price = request.getQueryString("pprice").flatMap(s => Try(s.toDouble).toOption)
    val manuf = request.getQueryString("pman").flatMap(s => Try(s.toString).toOption)
    if (name.exists(_.trim.nonEmpty) && price.isDefined && manuf.exists(_.trim.nonEmpty))
      {
        var newitem= new Products(pname = name.get, pprice = price.get, pman = manuf.get)
        prodlist+=newitem
        Ok("Dodano element do listy")
      }
    else
      BadGateway(s"Elementu nie udało się dodać do listy... $name $price $manuf")
  }
  def pmodify() = Action { implicit request =>
    val name = request.getQueryString("pname").flatMap(s => Try(s.toString).toOption)
    val price = request.getQueryString("pprice").flatMap(s => Try(s.toDouble).toOption)
    val manuf = request.getQueryString("pman").flatMap(s => Try(s.toString).toOption)
    val index = request.getQueryString("pmid").flatMap(s => Try(s.toInt).toOption)
    var ok=false
    if (index.isDefined) {
      ok = listsize(index.get)
    }
    if (ok==true)
      {
        if (name.exists(_.trim.nonEmpty)) {
          prodlist(index.get).pname=name.get
        }
        if (price.isDefined)
          {
            prodlist(index.get).pprice=price.get
          }
        if (manuf.exists(_.trim.nonEmpty))
          {
            prodlist(index.get).pman=manuf.get
          }
        val product = index.flatMap(i => prodlist.lift(i))
        Ok(views.html.pgetone(product))
      }
    else
      BadGateway(s"Elementu nie udało się zmodyfikować")
  }
}