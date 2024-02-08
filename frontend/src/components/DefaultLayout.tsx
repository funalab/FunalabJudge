import { ChakraProvider, Box, Flex, theme } from "@chakra-ui/react";
import { TopHeader } from "./TopHeader";
import { SideMenu } from "./SideMenu";
import React, { ReactNode } from 'react';

interface DefaultLayoutProps {
  children: ReactNode;
}

export const DefaultLayout: React.FC<DefaultLayoutProps> = ({ children }) => {
  return (
    <Flex w="100vw" h="100wh">
      <TopHeader />
      <Box mt="100px">
        <Flex>
            <SideMenu />
            <Box w="70vw">
              {children}
            </Box>
        </Flex>
      </Box>
    </Flex>
  );
};

export default DefaultLayout