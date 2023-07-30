import {useCallback, useState} from 'react';
import './App.css';
import {Greet} from "../wailsjs/go/main/App";

function App() {
    // const results = useState<string[]>([])
    const ref = useCallback((node: HTMLDivElement | null) => {
        node?.focus()
    }, [])

    const results = [
        "test1",
        "test2",
        "test3",
    ]

    return (
        <div id="App">
            <div id="word">
              <input type="text" ref={ref} />
            </div>
            <div id="results">
                {results.map(result => (
                    <div className="row">
                        <div>{result}</div>
                    </div>
                ))}
            </div>
        </div>
    )
}

export default App
