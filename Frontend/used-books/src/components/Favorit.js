import Link from 'react';
import React from 'react';
import Navbar from './Navbar';

function Favorit() {
	return (
        <div className="antialiased flex flex-col h-screen justify-center items-center bg-book">
            <Navbar />
            <div className="grid grid-cols-4 gap-24 my-10">
                {/* CRAZY LOVE */}
                <div className="border-2 rounded-xl items-center bg-white px-10 py-5 gap-6">
                    <div>
                        <img src='/crazy-love.png' className="" />
                    </div>
                    <div>
                        <div className='text-xl font-bold'>Crazy Love</div>
                        <div className='italic text-lg'>Francis Chan</div>
                    </div>
                    <div className='bg-primary mt-3 w-full'>
                        <div className='text-white text-lg py-3'>Rp. 170.000</div>
                    </div>
                </div>

                {/* YOUR HEART */}
                <div className="border-2 rounded-xl items-center bg-white px-10 py-5 gap-6">
                    <div>
                        <img src='/your-heart.png' className="" />
                    </div>
                    <div>
                        <div className='text-xl font-bold'>Your Heart Is The Sea</div>
                        <div className='italic text-lg'>Nikita Gill</div>
                    </div>
                    <div className='bg-primary mt-3 w-full'>
                        <div className='text-white text-lg py-3'>Rp. 120.000</div>
                    </div>
                </div>

                {/* MILK AND HONEY */}
                <div className="border-2 rounded-xl items-center bg-white px-10 py-5 gap-6 ">
                    <div>
                        <img src='/milk-honey.png' className="" />
                    </div>
                    <div>
                        <div className='text-xl font-bold'>Milk and Honey</div>
                        <div className='italic text-lg'>Rupi Kaur</div>
                    </div>
                    <div className='bg-primary mt-3 w-full'>
                        <div className='text-white text-lg py-3'>Rp. 160.000</div>
                    </div>
                </div>

                {/* STUPORE E TREMORI */}
                <div className="border-2 rounded-xl items-center bg-white px-10 py-5 gap-6 ">
                    <div>
                        <img src='/stupore.png' className="" />
                    </div>
                    <div>
                        <div className='text-xl font-bold'>Stupore E Tremori</div>
                        <div className='italic text-lg'>Amelie Nothomb</div>
                    </div>
                    <div className='bg-primary mt-3 w-full'>
                        <div className='text-white text-lg py-3'>Rp. 109.000</div>
                    </div>
                </div>

                {/* psycho-money */}
                <div className="border-2 rounded-xl items-center bg-white px-10 py-5 ">
                    <div>
                        <img src='/crazy-love.png' className="" />
                    </div>
                    <div>
                        <div className='text-xl font-bold'>The Psychology of Money</div>
                        <div className='italic text-lg'>Morgan Housel</div>
                    </div>
                    <div className='bg-primary mt-3 w-full'>
                        <div className='text-white text-lg py-3'>Rp. 120.000</div>
                    </div>
                </div>

                {/* HOW INNOVATION WORKS */}
                <div className="border-2 rounded-xl items-center bg-white px-10 py-5 ">
                    <div>
                        <img src='/innovation-works.png' className="" />
                    </div>
                    <div>
                        <div className='text-xl font-bold'>How Innovation Works</div>
                        <div className='italic text-lg'>Matt Ridley</div>
                    </div>
                    <div className='bg-primary mt-3 w-full'>
                        <div className='text-white text-lg py-3'>Rp. 160.000</div>
                    </div>
                </div>

                 {/* CRAZY LOVE */}
                 <div className="border-2 rounded-xl items-center bg-white px-10 py-5 gap-6">
                    <div>
                        <img src='/crazy-love.png' className="" />
                    </div>
                    <div>
                        <div className='text-xl font-bold'>Crazy Love</div>
                        <div className='italic text-lg'>Francis Chan</div>
                    </div>
                    <div className='bg-primary mt-3 w-full'>
                        <div className='text-white text-lg py-3'>Rp. 170.000</div>
                    </div>
                </div>

                {/* YOUR HEART */}
                <div className="border-2 rounded-xl items-center bg-white px-10 py-5 gap-6">
                    <div>
                        <img src='/your-heart.png' className="" />
                    </div>
                    <div>
                        <div className='text-xl font-bold'>Your Heart Is The Sea</div>
                        <div className='italic text-lg'>Nikita Gill</div>
                    </div>
                    <div className='bg-primary mt-3 w-full'>
                        <div className='text-white text-lg py-3'>Rp. 120.000</div>
                    </div>
                </div>

                {/* MILK AND HONEY */}
                <div className="border-2 rounded-xl items-center bg-white px-10 py-5 gap-6 ">
                    <div>
                        <img src='/milk-honey.png' className="" />
                    </div>
                    <div>
                        <div className='text-xl font-bold'>Milk and Honey</div>
                        <div className='italic text-lg'>Rupi Kaur</div>
                    </div>
                    <div className='bg-primary mt-3 w-full'>
                        <div className='text-white text-lg py-3'>Rp. 160.000</div>
                    </div>
                </div>

                {/* STUPORE E TREMORI */}
                <div className="border-2 rounded-xl items-center bg-white px-10 py-5 gap-6 ">
                    <div>
                        <img src='/stupore.png' className="" />
                    </div>
                    <div>
                        <div className='text-xl font-bold'>Stupore E Tremori</div>
                        <div className='italic text-lg'>Amelie Nothomb</div>
                    </div>
                    <div className='bg-primary mt-3 w-full'>
                        <div className='text-white text-lg py-3'>Rp. 109.000</div>
                    </div>
                </div>

                {/* psycho-money */}
                <div className="border-2 rounded-xl items-center bg-white px-10 py-5 ">
                    <div>
                        <img src='/crazy-love.png' className="" />
                    </div>
                    <div>
                        <div className='text-xl font-bold'>The Psychology of Money</div>
                        <div className='italic text-lg'>Morgan Housel</div>
                    </div>
                    <div className='bg-primary mt-3 w-full'>
                        <div className='text-white text-lg py-3'>Rp. 120.000</div>
                    </div>
                </div>

                {/* HOW INNOVATION WORKS */}
                <div className="border-2 rounded-xl items-center bg-white px-10 py-5 ">
                    <div>
                        <img src='/innovation-works.png' className="" />
                    </div>
                    <div>
                        <div className='text-xl font-bold'>How Innovation Works</div>
                        <div className='italic text-lg'>Matt Ridley</div>
                    </div>
                    <div className='bg-primary mt-3 w-full'>
                        <div className='text-white text-lg py-3'>Rp. 160.000</div>
                    </div>
                </div>
            </div>
        </div>
    );
}

export default Favorit;
