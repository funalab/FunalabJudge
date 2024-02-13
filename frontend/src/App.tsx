import { BrowserRouter, Routes, Route } from "react-router-dom";

import { AxiosClientProvider } from "./providers/AxiosClientProvider";

import Login from "./routes/Login";
import Account from "./routes/Account";
import Dashboard from "./routes/Dashboard";
import Message from "./routes/Message";
import Schedule from "./routes/Schedule";
import AssignmentPage from "./routes/AssignmentPage";
import ResultQueuePage from "./routes/ResultQueuePage";
import SubmitDetailsPage from "./routes/SubmitDetailsPage";
import NotFound from "./routes/NotFound";

const App: React.FC = () => {
  return (
    <BrowserRouter>
      <AxiosClientProvider>
        <Routes>
          <Route path="/login" element={<Login />} />
          <Route path="*" element={<NotFound />} />
          <Route path="/:userName">
            <Route path="account" element={<Account />} />
            <Route path="dashboard" element={<Dashboard />} />
            <Route path="message" element={<Message />} />
            <Route path="schedule" element={<Schedule />} />
            <Route path="assignmentInfo/:id" element={<AssignmentPage />} />
            <Route path="results" element={<ResultQueuePage />} />
            <Route path="submission/:submitId" element={<SubmitDetailsPage />} />
          </Route>
        </Routes>
      </AxiosClientProvider>
    </BrowserRouter>
  );
};

export default App;
