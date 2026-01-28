import { useState } from "react";
import api from "../api/client";
import { useNavigate } from "react-router-dom";

export default function Login() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const nav = useNavigate();

  const login = async () => {
    const res = await api.post("/auth/login", { email, password });
    localStorage.setItem("token", res.data.token);
    nav("/dashboard");
  };

  return (
    <div className="h-screen flex items-center justify-center">
      <div className="w-80 p-6 border rounded">
        <h1 className="text-xl mb-4">Login</h1>
        <input className="input" placeholder="Email" onChange={e => setEmail(e.target.value)} />
        <input className="input mt-2" type="password" placeholder="Password" onChange={e => setPassword(e.target.value)} />
        <button className="mt-4 w-full bg-black text-white p-2" onClick={login}>
          Login
        </button>
      </div>
    </div>
  );
}
