// ルーティング設定に必要なものをimport
import { BrowserRouter, Routes, Route } from "react-router-dom";
// ルーティング先の画面コンポーネントをimport
import { Login } from "./routes/Login";
import axios from "axios";
import Account from "./routes/Account";
import Dashboard from "./routes/Dashboard";
import Message from "./routes/Message";
import Schedule from "./routes/Schedule";
import AssignmentPage from "./routes/AssignmentPage";

axios.defaults.baseURL = 'http://localhost:3000';

const App: React.FC = () => {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Dashboard />} />
        <Route path="/login" element={<Login />} />
        <Route path="/account" element={<Account />} />
        <Route path="/dashboard" element={<Dashboard />} />
        <Route path="/message" element={<Message />} />
        <Route path="/schedule" element={<Schedule />} />
        <Route path="/assignmentInfo/:id" element={<AssignmentPage />} />
      </Routes>
    </BrowserRouter>
  );
};

export default App;
