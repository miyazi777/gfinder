import {useCallback, useEffect, useState} from 'react';
import './App.css';
import {Exec, GetInitialList, Quit, Search} from "../wailsjs/go/main/App";
import {main} from "../wailsjs/go/models";

function App() {
    const [selected, setSelected] = useState(0);
    const [word, setWord] = useState("");
    const [results, setResults] = useState<main.InnerResource[]>([]);
    const [isShowConfirmModal, setIsShowConfirmModal] = useState(false);
    const [confirmSelected, setConfirmSelected] = useState(0);

    const ref = useCallback((node: HTMLDivElement | null) => {
        node?.focus()
    }, [])

    useEffect(() => {
        GetInitialList().then((list) => {
            setResults(list);
        });
    }, [])

    const handleEnter = () => {
        if (results[selected].confirm_dialog) {
            setIsShowConfirmModal(true);
            return;
        }
        Exec(results[selected]);
    }

    const handleConfirmEnter = () => {
        if (confirmItems[confirmSelected] === 'Yes') {
            Exec(results[selected]);
        }
        setIsShowConfirmModal(false); 
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
                isShowConfirmModal ?
                    setConfirmSelected((confirmSelected - 1) < 0 ? confirmItems.length - 1 : confirmSelected - 1) :
                    setSelected((selected - 1) < 0 ? results.length - 1 : selected - 1);
                break;
            case 'Down': 
            case 'ArrowDown': 
                isShowConfirmModal ?
                    setConfirmSelected((confirmSelected - 1) >= confirmItems.length ? 0 : confirmSelected + 1) :
                    setSelected((selected + 1) >= results.length ? 0 : selected + 1);
                break;
            case 'Enter':
                isShowConfirmModal ? handleConfirmEnter() : handleEnter();
                break;
            case 'Escape':  
                handleQuit();
                break;
        }
    }

    const confirmItems = ['Yes', 'No'];

    return (
        <>
            <div id="App" onKeyDown={handleKeyDown} tabIndex={0}>
                <>
                    <div id="word">
                      <input type="text" ref={ref} value={word} onChange={e => handleChange(e.target.value)} />
                    </div>
                    {!isShowConfirmModal && (
                        <div id="results">
                            {results.map((result, index) => (
                                <div key={index} className={`row ${index === selected ? "selected" : ""}`}>
                                    <div>
                                        <div className="name">{result.name}</div>
                                        <div className="info">{result.info}</div>
                                    </div>
                                    <div className="tag">[{result.tag}]</div>
                                </div>
                            ))}
                        </div>
                    )}
                    {isShowConfirmModal && (
                        <div id="results">
                            <div>{results[selected].name}</div>
                            {confirmItems.map((item, index) => (
                                <div key={index} className={`row ${index === confirmSelected ? "selected" : ""}`}>
                                    <div>{item}</div>
                                </div>
                            ))}
                        </div>
                    )}
                </>
            </div>
        </>
    )
}

export default App
