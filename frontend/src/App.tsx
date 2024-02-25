import { BrowserRouter, Routes, Route } from "react-router-dom";

import { AxiosClientProvider } from "./providers/AxiosClientProvider";

import Login from "./routes/Login";
import AccountPage from "./routes/Account";
import DashboardPage from "./routes/Dashboard";
import MessagePage from "./routes/Message";
import SchedulePage from "./routes/Schedule";
import ProblemPage from "./routes/Problem";
import ResultsPage from "./routes/Results";
import SubmissionPage from "./routes/Submission";
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
            <Route path="account" element={<RouteAuthGuard component={<AccountPage />} pageType={PageType.Private} />} />
            <Route path="dashboard" element={<RouteAuthGuard component={<DashboardPage />} pageType={PageType.Private} />} />
            <Route path="message" element={<RouteAuthGuard component={<MessagePage />} pageType={PageType.Private} />} />
            <Route path="schedule" element={<RouteAuthGuard component={<SchedulePage />} pageType={PageType.Private} />} />
            <Route path="problem/:problemId" element={<RouteAuthGuard component={<ProblemPage />} pageType={PageType.Private} />} />
            <Route path="results" element={<RouteAuthGuard component={<ResultsPage />} pageType={PageType.Private} />} />
            <Route path="submission/:submissionId" element={<RouteAuthGuard component={<SubmissionPage />} pageType={PageType.Private} />} />
          </Route>
          <Route path="*" element={<NotFound />} />
        </Routes>
      </AxiosClientProvider>
    </BrowserRouter>
  );
};
export default App;
