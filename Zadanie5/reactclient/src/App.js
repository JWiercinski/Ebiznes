import './App.css';
import React from "react";
import {Routes} from "react-router-dom";
import {Route, BrowserRouter} from "react-router-dom";
import Hub from "./Hub";
import Products from "./Products";
import DefaultApp from "./DefaultApp";
import Payments from "./Payments";

function App() {
    return (
        <div className="App">
            <div className="content">
                <BrowserRouter>
                    <Routes>
                        <Route path="/" element={<Hub/>}></Route>
                        <Route path="/products" element={<Products/>}></Route>
                        <Route path="/payments" element={<Payments/>}></Route>
                        <Route path="/x2j23fv" element={<DefaultApp/>}></Route>
                    </Routes>
                </BrowserRouter>
            </div>
        </div>
    );
};

export default App;
