import {useCallback, useEffect, useState} from 'react';
import './App.css';
import {Exec, GetInitialList, Greet, Quit, Search} from "../wailsjs/go/main/App";

function App() {
    const [selected, setSelected] = useState(0);
    const [word, setWord] = useState("");
    const [results, setResults] = useState<string[]>([]);

    const ref = useCallback((node: HTMLDivElement | null) => {
        node?.focus()
    }, [])

    useEffect(() => {
        GetInitialList().then((list: string[]) => {
            setResults(list);
        });
    }, [])

    const handleEnter = () => {
        Exec(results[selected]);
    }

    const handleQuit = () => {
        Quit();
    }

    const handleChange = (value: string) => {
        setWord(value);
        Search(value).then(searchedResults => {
            setResults(searchedResults)
        })

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
              <input type="text" ref={ref} value={word} onChange={e => handleChange(e.target.value)} />
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
