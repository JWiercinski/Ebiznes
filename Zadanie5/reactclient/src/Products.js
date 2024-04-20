import React from "react";
import axios from "axios";

const Products = () => {
    const [products, setProducts] = React.useState([]);

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
            <h1>Produkty</h1>
            {products.map((product, index) => (
                <div key={index}>
                    <h2>{product.NAME}</h2>
                    <p>{product.DESC}</p>
                    <p>Twórca: {product.DEV}</p>
                    <p>Cena: {product.PRICE} PLN</p>
                </div>
            ))}
        </div>
    );
};
export default Products