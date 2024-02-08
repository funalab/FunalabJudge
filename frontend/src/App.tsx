import { ChakraProvider, theme } from "@chakra-ui/react";

// ルーティング設定に必要なものをimport
import { BrowserRouter, Routes, Route } from "react-router-dom";

// ルーティング先の画面コンポーネントをimport
import { Login } from "./routes/Login";
import { Account } from "./routes/Account";
import { Dashboard } from "./routes/Dashboard";
import { Message } from "./routes/Message";
import { Schedule } from "./routes/Schedule";

export const App = () => {
  return (
    <ChakraProvider theme={theme}>
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<Dashboard />} />
          <Route path="/login" element={<Login />}/> 
          <Route path="/account" element={<Account />} />
          <Route path="/dashboard" element={<Dashboard />} />
          <Route path="/message" element={<Message />} />
          <Route path="/schedule" element={<Schedule />} />
        </Routes>
      </BrowserRouter>
    </ChakraProvider>
  );
};

export default App
