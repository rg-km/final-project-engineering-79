import logo from './logo.svg';
import './App.css';
import Link from 'react';
import React from 'react';

function App() {
	return (
        <div className="antialiased flex h-screen justify-center items-center bg-book">
            <div className="flex w-1/2 items-center justify-center">
                <div className="w-4/5 items-center justify-center">
                    <div className="grid gap-2 my-4">
                        <div className= "justify-center items-center text-center">
                            <img className="w-40 " src="/logo.png" />
                            <div className="text-2xl tracking-wide">
                                Masuk
                            </div>
                        </div>
                        
                    </div>
                    

                
                </div>
            </div>
        </div>
    );
}

export default App;
