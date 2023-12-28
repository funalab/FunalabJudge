import { ChakraProvider, Box, Flex, theme } from "@chakra-ui/react";
import { TopHeader } from "./components/TopHeader";
import { SideMenu } from "./components/SideMenu";

// ルーティング設定に必要なものをimport
import { BrowserRouter, Routes, Route } from "react-router-dom";

// ルーティング先の画面コンポーネントをimport
import { Account } from "./routes/Account";
import { Dashboard } from "./routes/Dashboard";
import { Message } from "./routes/Message";
import { Schedule } from "./routes/Schedule";

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
                  <Route path="/" element={<Dashboard />} />
                  <Route path="/account" element={<Account />} />
                  <Route path="/dashboard" element={<Dashboard />} />
                  <Route path="/message" element={<Message />} />
                  <Route path="/schedule" element={<Schedule />} />
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
