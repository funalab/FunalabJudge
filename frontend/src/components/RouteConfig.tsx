import React from "react";
import { BrowserRouter, Routes, Route } from "react-router-dom";

// ルーティング先の画面コンポーネントをimport
import { Login } from "../routes/Login";
import { Account } from "../routes/Account";
import { Dashboard } from "../routes/Dashboard";
import { Message } from "../routes/Message";
import { Schedule } from "../routes/Schedule";
import { NotFound } from "../routes/NotFound";

import { RouteAuthGuard } from "./RouteAuthGuard";
import { PageType } from "../types/PageTypes";

export const RouterConfig:React.FC = () => {
  // <Route path="/dashboard" element={<RouteAuthGuard component={<Dashboard />} pagetype={private} redirect="/login" />} />
  // みたいにしたい
  return (
  <>
    <BrowserRouter>
      <Routes>
        <Route path="/:userName">
          <Route path="dashboard" element={<Dashboard />} />
          <Route path="account" element={<Account />} />
          <Route path="message" element={<Message />} />
          <Route path="schedule" element={<Schedule />} />
        </Route>
        <Route path="/login" element={<Login />}/>
        <Route path="*" element={<NotFound />} />
      </Routes>
    </BrowserRouter>
  </>
  );
}