import React from "react";
import { Providers } from "./components/Providers"
import { RouterConfig } from "./components/RouteConfig";

export const App: React.FC = () => {
  return (
    <Providers>
      <RouterConfig />
    </Providers>
  );
};

export default App
