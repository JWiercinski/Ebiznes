import React, {useState} from "react";
import axios from "axios";
import {BasketContext} from "./Basket";
import {Link} from "react-router-dom";

function Payments(){
    const [user, setUser] = useState("")
    const [method, setMethod] = useState("")
    //const [price, setPrice]=useState("")
    const [info, setInfo]=useState("")
    const { basket} = React.useContext(BasketContext);
    const total = basket.reduce((accumulator, item) => {
        return accumulator + (item.product.PRICE * item.quantity);
    }, 0);
    const sendonsubmit = async (event) => {
        event.preventDefault()
        if (!method) {
            alert('Wybierz właściwą metodę płatności');
            return;
        }
        const data = {
            USER: user,
            METHOD: method,
            AMOUNT: total
        }
        try {
            const response = await axios.post("https://jwgoserver.azurewebsites.net/payment", data)
            setInfo(`Sukces!`)
            for (let i = 0; i < basket.length; i++) {
                const item = basket[i];
                const data2 ={
                    PAYID: parseInt(response.data),
                    GAMEID: item.product.ID,
                    QUANTITY: item.quantity
                }
                try {
                    await axios.post("https://jwgoserver.azurewebsites.net/basket", data2)
                }
                catch (e) {
                    setInfo(e)
                }
            }
            localStorage.removeItem("basket")
            window.location.reload()
        }
        catch (error){
            setInfo(`Mamy problem, nie mogliśmy zrealizować płatności. ${error.response.data}`)
        }
    }
    return (
        <div className="Payments">
            <Link to="/"><button>Strona główna</button></Link>
            <Link to="/products"><button>Kontynuuj zakupy</button></Link>
            <Link to="/basket"><button>Wróć do koszyka</button></Link>
            <h1>Dokonaj płatności</h1>
            <form onSubmit={sendonsubmit}>
                <p>Nazwa Użytkownika</p>
                <input type="text" value={user} onChange={(e)=>setUser(e.target.value)} placeholder="Nazwa użytkownika" required></input>
                <p>Metoda Płatności</p>
                <select value={method} onChange={(e) => setMethod(e.target.value)}>
                    <option value="" disabled>Wybierz metodę</option>
                    <option value="CARD">Karta płatnicza</option>
                    <option value="BANKTRANSFER">Przelew bankowy</option>
                    <option value="PAYPAL">PayPal</option>
                </select>
                <h3>Kwota: {total} PLN</h3>
                <button type="submit">Kontynuuj</button>
            </form>
            <h2>{info}</h2>
        </div>
    )
}

export default Payments
