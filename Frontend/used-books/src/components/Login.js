import React, { useState } from "react";
import ReactDOM from "react-dom";

function Login() {
	const [errorMessages, setErrorMessages] = useState({});
	const [isSubmitted, setIsSubmitted] = useState(false);

	const database = [
		{
			username: "user1",
			password: "pass1"
		},
		{
			username: "user2",
			password: "pass2"
		}
	];
	
	const errors = {
		uname: "invalid username",
		pass: "invalid password"
	};

	const handleSubmit = (event) => {
		event.preventDefault();
	
		var { uname, pass } = document.forms[0];
		const userData = database.find((user) => user.username === uname.value);
	
		// Compare user info
		if (userData) {
			if (userData.password !== pass.value) {
			// Invalid password
			setErrorMessages({ name: "pass", message: errors.pass });
			} else {
				setIsSubmitted(true);
			}
		} else {
		  // Username not found
			setErrorMessages({ name: "uname", message: errors.uname });
		}
	};
	
	// Generate JSX code for error message
	const renderErrorMessage = (name) =>
    name === errorMessages.name && (
		<div className="error">{errorMessages.message}</div>
    );

	// JSX code for login form
	const renderForm = (
		<div className="form">
			<form onSubmit={handleSubmit}>
				<div className="input-container">
					{/* <label>Username </label> */}
					{/* <input type="text" name="uname" required /> */}
					{renderErrorMessage("uname")}
				</div>
				<div className="input-container">
					{/* <label>Password </label> */}
					{/* <input type="password" name="pass" required /> */}
					{renderErrorMessage("pass")}
				</div>
				{/* <div className="button-container">
					<input type="submit" />
				</div> */}
			</form>
		</div>
	);

	return (
		<div className="antialiased flex h-screen justify-center items-center bg-book">
			<div className="flex w-1/2 items-center justify-center">
				<div className="w-4/5 items-center justify-center">
					<div className="grid gap-2 my-4">
						<div className= "justify-center items-center text-center">
							<img className="w-40 " src="/logo.png" />
							<div className="text-3xl tracking-wide font-bold">
								Masuk
							</div>
						</div>
						
					</div>
					
					{/* EMAIL  */}
					<div className="mb-4 block">
						<div className="relative block">
							<input
								type="email"
								name="email"
								className="placeholder:text-gray-400 block bg-white w-full border border-gray-300 rounded-xl py-4 px-4 shadow-md focus:outline-none focus:border-sky-500 focus:ring-sky-500 focus:ring-1 placeholder:text-lg"
								placeholder="Email"
							/>
						</div>
					</div>

					{/* PASSWORD  */}
					<div className="mb-4 block">
						<label className="relative block">
							<div className="absolute inset-y-0 left-0 flex items-center pl-4">
								<img src="/images/password.png" alt="" className="w-4" />
							</div>
							<input
								type="password"
								name="password"
								className="placeholder:text-gray-400 block bg-white w-full border border-gray-300 rounded-xl py-4 px-4 shadow-md focus:outline-none focus:border-sky-500  focus:ring-sky-500 focus:ring-1 placeholder:text-lg"
								placeholder="Password"
							/>
						</label>
					</div>		

					{/* LOGIN  */}
					<button className="flex items-center justify-center gap-3 w-full bg-primary font-bold px-2 py-4 hover:bg-opacity-90 text-white border-2 rounded-xl text-xl cursor-pointer my-6">
						<div className="">
							<img src="/images/landing/building.png" className="h-6" alt="" />
						</div>
						<div className="text-lg lg:text-xl">Login</div>
						{isSubmitted ? <div>User is successfully logged in</div> : renderForm}
					</button>		
							
					
					<div className="flex items-center justify-center text-lg font-bold">
						<div className="border-t-2" />
							<p>Sudah ada akun?</p>
						<div className="border-t-2 bg-red-200" />
					</div>

					{/* DAFTAR */}
					<button className="flex items-center justify-center gap-3 w-full bg-green-300 font-bold px-2 py-4 hover:bg-opacity-90 text-white border-2 rounded-xl text-xl cursor-pointer my-6">
						<div className="">
							<img src="/images/landing/building.png" className="h-6" alt="" />
						</div>
						<div className="text-lg lg:text-xl">Daftar</div>
					</button>
				</div>
			</div>
		</div>
	);
}

export default Login;
