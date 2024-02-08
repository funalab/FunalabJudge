import { ChakraProvider, Box, Flex, theme } from "@chakra-ui/react";
import { TopHeader } from "./components/TopHeader";
import { SideMenu } from "./components/SideMenu";

import { BrowserRouter, Routes, Route } from "react-router-dom";
import DashboardPage from "./routes/DashboardPage";
import AccountPage from "./routes/AccountPage";
import MessagePage from "./routes/MessagePage";
import SchedulePage from "./routes/SchedulePage";
import AssignmentPage from "./routes/AssignmentPage";


export const App = () => {
  return (
    <ChakraProvider theme={theme}>
      <Flex w="100vw" h="100wh">
        <BrowserRouter>
          <TopHeader />
        </BrowserRouter>
        <Box mt="100px">
          <Flex>
            <BrowserRouter>
              <SideMenu />
              <Box w="70vw">
                <Routes>
                  <Route path="/" element={<DashboardPage />} />
                  <Route path="/account" element={<AccountPage />} />
                  <Route path="/dashboard" element={<DashboardPage />} />
                  <Route path="/message" element={<MessagePage />} />
                  <Route path="/schedule" element={<SchedulePage />} />
                  <Route path="/assignment" element={<AssignmentPage />} />
                </Routes>
              </Box>
            </BrowserRouter>
          </Flex>
        </Box>
      </Flex>
    </ChakraProvider>
  );
};

export default App
