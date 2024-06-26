import React from 'react';
import {Link} from "react-router-dom";
export const BasketContext = React.createContext();
export const BasketProvider = ({ children }) => {
    const [basket, setBasket] = React.useState(() => {
        // Try to get the basket from localStorage
        const savedBasket = localStorage.getItem('basket');
        if (savedBasket) {
            return JSON.parse(savedBasket);
        }
        return [];
    });
    React.useEffect(() => {
        localStorage.setItem('basket', JSON.stringify(basket));
    }, [basket]);
    return (
        <BasketContext.Provider value={{ basket, setBasket }}>
            {children}
        </BasketContext.Provider>
    );
};

function Basket () {
    const { basket } = React.useContext(BasketContext);
    const clearBasket = () => {
        localStorage.removeItem("basket")
        window.location.reload()
    }
    const total = basket.reduce((accumulator, item) => {
        return accumulator + (item.product.PRICE * item.quantity);
    }, 0);

    return (
        <div>
            <Link to="/"><button>Strona główna</button></Link>
            <Link to="/products"><button>Kontynuuj zakupy</button></Link>
            <h1>Your Basket:</h1>
            {basket.map((item, index) => (
                <div key={index}>
                    <h2>{item.product.NAME}</h2>
                    <p>Cena jednostkowa: {item.product.PRICE} PLN</p>
                    <p>Ilość: {item.quantity}</p>
                </div>
            ))}
           <h2>Podsuma: {total} PLN</h2>
            <button style={{backgroundColor:"darkred", color:"white"}} onClick={() => clearBasket()}>Wyczyść koszyk</button>
            <Link to="/payments"><button>Przejdź do płatności</button></Link>
            <p></p>
        </div>
    );
};

export default Basket