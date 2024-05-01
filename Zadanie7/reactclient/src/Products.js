import React from "react";
import axios from "axios";
import {BasketContext} from "./Basket";
import {Link} from "react-router-dom";

const Products = () => {
    const [products, setProducts] = React.useState([]);
    const { basket, setBasket } = React.useContext(BasketContext);

    const addToBasket = (product) => {
        const existingProductIndex = basket.findIndex(item => item.product.ID === product.ID);

        if (existingProductIndex !== -1) {
            const updatedBasket = [...basket];
            updatedBasket[existingProductIndex] = {
                ...updatedBasket[existingProductIndex],
                quantity: updatedBasket[existingProductIndex].quantity + 1
            };
            setBasket(updatedBasket);
        } else {
            setBasket([...basket, { product, quantity: 1 }]);
        }
    };

    React.useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await axios.get('http://localhost:22222/products');
                setProducts(response.data);
            } catch (error) {
                console.error('Error fetching data: ', error);
            }
        };
        fetchData();
    }, []);

    return (
        <div>
            <Link to="/"><button>Strona główna</button></Link>
            <Link to="/basket"><button>Przejdź do koszyka</button></Link>
            <h1>Produkty</h1>
            {products.map((product) => (
                <div key={product.ID}>
                    <h2>{product.NAME}</h2>
                    <p>{product.DESC}</p>
                    <p>Twórca: {product.DEV}</p>
                    <p>Cena: {product.PRICE} PLN</p>
                    <button onClick={()=> addToBasket(product)}>Dodaj do koszyka</button>
                </div>
            ))}
        </div>
    );
};
export default Products