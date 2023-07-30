import {useCallback, useState} from 'react';
import './App.css';
import {Greet, Quit} from "../wailsjs/go/main/App";

function App() {
    const [selected, setSelected] = useState(0);

    const ref = useCallback((node: HTMLDivElement | null) => {
        node?.focus()
    }, [])

    const results = [
        "test1",
        "test2",
        "test3",
        "test4",
        "test5",
        "test6",
        "test7",
    ];

    const handleEnter = () => {
        console.log('selected', results[selected]);
        // Exec(results[selected]).then();
    }

    const handleQuit = () => {
        Quit();
    }


    const handleKeyDown = (e: React.KeyboardEvent<HTMLDivElement>) => {
        const key = e.code;

        switch (key) {
            case 'Up':
            case 'ArrowUp':
                setSelected((selected - 1) < 0 ? results.length - 1 : selected - 1);
                break;
            case 'Down': 
            case 'ArrowDown': 
                setSelected((selected + 1) >= results.length ? 0 : selected + 1);
                break;
            case 'Enter':
                handleEnter();
                break;
            case 'Escape':  
                handleQuit();
                break;
        }
    }

    return (
        <div id="App" onKeyDown={handleKeyDown}>
            <div id="word">
              <input type="text" ref={ref} />
            </div>
            <div id="results">
                {results.map((result, index) => (
                    <div key={index} className={`row ${index === selected ? "selected" : ""}`}>
                        <div>{result}</div>
                    </div>
                ))}
            </div>
        </div>
    )
}

export default App
