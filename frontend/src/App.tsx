import { BrowserRouter, Routes, Route, Navigate } from "react-router-dom";
import { useEffect } from "react";

import { AxiosClientProvider } from "./providers/AxiosClientProvider";

import Login from "./routes/Login";
import AccountPage from "./routes/Account";
import DashboardPage from "./routes/Dashboard";
import SchedulePage from "./routes/PetitCoder";
import ProblemPage from "./routes/Problem";
import ResultsPage from "./routes/Results";
import SubmissionPage from "./routes/Submission";
import B3ResultsPage from "./routes/B3Results";
import NotFound from "./routes/NotFound";
import { RouteAuthGuard } from "./providers/RouteAuthGuard";
import { PageType } from "./types/PageTypes";

const App: React.FC = () => {
  useEffect(() => {
    document.title = 'Funalab Judge';
  }, []);
  return (
    <BrowserRouter>
      <AxiosClientProvider>
        <Routes>
          <Route path="/login" element={<Login />} />

          {/* 自分のコンテンツのみ表示可能なページ */}
          <Route path="/dashboard" element={<RouteAuthGuard component={<DashboardPage />} pageType={PageType.Public} />} />
          <Route path="/account" element={<RouteAuthGuard component={<AccountPage />} pageType={PageType.Public} />} />
          {/* 全員共通のコンテンツを持つページ */}
          <Route path="/problem/:problemId" element={<RouteAuthGuard component={<ProblemPage />} pageType={PageType.Public} />} />
          <Route path="/all_results" element={<RouteAuthGuard component={<B3ResultsPage />} pageType={PageType.Public} />} />
          <Route path="/petit_coder" element={<RouteAuthGuard component={<SchedulePage />} pageType={PageType.Public} />} />
          {/* 上級生はB3のコンテンツも閲覧可能なページ */}
          <Route path="/results/:userName" element={<RouteAuthGuard component={<ResultsPage />} pageType={PageType.Private} />} />
          {/* resultページにしかリンクが存在しないページ */}
          <Route path="/submission/:submissionId" element={<RouteAuthGuard component={<SubmissionPage />} pageType={PageType.Public} />} />

          <Route path="/" element={<Navigate to="/login" />} />
          <Route path="/*" element={<NotFound />} />
        </Routes>
      </AxiosClientProvider>
    </BrowserRouter>
  );
};
export default App;
