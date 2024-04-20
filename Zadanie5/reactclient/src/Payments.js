import {useState} from "react";
import axios from "axios";

function Payments(){
    const [user, setUser] = useState("")
    const [method, setMethod] = useState("")
    const [price, setPrice]=useState("")
    const [info, setInfo]=useState("")
    const sendonsubmit = async (event) => {
        event.preventDefault()
        if (!method) {
            alert('Wybierz właściwą metodę płatności');
            return;
        }
        const data = {
            USER: user,
            METHOD: method,
            AMOUNT: parseFloat(price)
        }
        try {
            const response = await axios.post("http://localhost:22222/payment", data)
            setInfo(`Sukces!`)
        }
        catch (error){
            setInfo(`Mamy problem, nie mogliśmy zrealizować płatności. ${error.response.data}`)
        }
    }
    return (
        <div className="Payments">
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
                <p>Kwota</p>
                <input type="number" step="0.01" value={price} onChange={(e)=>setPrice(e.target.value)} placeholder="Cena do zapłaty" required></input>
                <p></p>
                <button type="submit">Kontynuuj</button>
            </form>
            <h2>{info}</h2>
        </div>
    )
}

export default Payments
/*
type payment struct {
    USER   string  `json:"USER"`
    METHOD string  `json:"METHOD"`
    AMOUNT float64 `json:"AMOUNT"`
}
*/