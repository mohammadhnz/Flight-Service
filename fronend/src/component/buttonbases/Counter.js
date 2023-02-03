import {useState} from 'react';
import {Input} from "@mui/material";


function Counter() {
    let [count, setCount] = useState(0)

    const increment = () => {
        setCount(count + 1)
    };

    const decrement = () => {
        setCount(count - 1)
    };

    return (
        <div className="counter">
            <Input>{count}</Input>
            <button onClick={increment}>+</button>
            <button onClick={decrement}>-</button>
        </div>
    );
}

export default Counter;
