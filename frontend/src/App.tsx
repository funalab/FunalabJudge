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
import { RouteAuthGuard } from "./providers/RouteAuthGuard";
import { PageType } from "./types/PageTypes";

const App: React.FC = () => {
  return (
    <BrowserRouter>
      <AxiosClientProvider>
        <Routes>
          <Route path="/login" element={<Login />} />
          <Route path="/:userName">
            <Route path="account" element={<RouteAuthGuard component={<Account />} pageType={PageType.Private} />} />
            <Route path="dashboard" element={<RouteAuthGuard component={<Dashboard />} pageType={PageType.Private} />} />
            <Route path="message" element={<RouteAuthGuard component={<Message />} pageType={PageType.Private} />} />
            <Route path="schedule" element={<RouteAuthGuard component={<Schedule />} pageType={PageType.Private} />} />
            <Route path="assignmentInfo/:problemId" element={<RouteAuthGuard component={<AssignmentPage />} pageType={PageType.Private} />} />
            <Route path="results/:problemId" element={<RouteAuthGuard component={<ResultQueuePage />} pageType={PageType.Private} />} />
            <Route path="submission/:submissionId" element={<RouteAuthGuard component={<SubmitDetailsPage />} pageType={PageType.Private} />} />
          </Route>
          <Route path="*" element={<NotFound />} />
        </Routes>
      </AxiosClientProvider>
    </BrowserRouter>
  );
};
export default App;
