import {Link} from "react-router-dom";
import React from "react";

function Hub() {
    return (
        <div className="Hub">
            <h1>Reaction Nation</h1>
            <h2>Wszystko co najlepsze i najgorsze - wszystko w jednym miejscu</h2>
            <Link to="/products"><button>Przejdź do produktów</button></Link>
            <Link to="/payments"><button>Przejdź do płatności</button></Link>
        </div>
    );
};

export default Hub