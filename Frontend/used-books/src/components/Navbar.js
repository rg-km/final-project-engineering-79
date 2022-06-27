import React from "react";
import Image from "react";

function Navbar(){
    return (
        <div aria-label="navbar" className="w-11/12">
            <header className="w-11/12 mx-auto grid grid-cols-12 items-center py-8 my-10 text-lg">
				<div className="col-span-1 lg:col-span-2">
					<div className="w-40">
						<img src="/logo.png" alt="" />
					</div>
				</div>
				<div className="hidden lg:flex col-span-5 gap-5 text-3xl font-bold">
					Wishlist
				</div>
				<div className="col-span-10 lg:col-span-5 justify-end flex gap-8 items-center">
					<div className="hidden lg:flex">
						<img src="/images/lowongan-detail/notification-gray.png" alt="" className="w-6" />
					</div>
					<div className="hidden lg:flex items-center gap-10">
						<div className=" border-4 rounded-full border-white overflow-hidden">
							<img src="/cart.png" alt="" className="w-14" />
						</div>
						<div>
							<img src="/love.png" alt="" className="w-12" />
						</div>
                        <div>
							<img src="/profile.png" alt="" className="w-20" />
						</div>
					</div>
				</div>
			</header>
        </div>
    );
}

export default Navbar;