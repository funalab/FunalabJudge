import { Box, Button, Icon } from "@chakra-ui/react";
import {
  MdDashboard,
  MdCalendarMonth,
  MdAccountBox,
  MdAssignment,
} from "react-icons/md";
import { useNavigate, useParams } from "react-router-dom";

export const SideMenu = () => {
  const navigate = useNavigate();
  const { userName } = useParams();
  const sideMenuItems = [
    {
      name: "Dashboard",
      icon: MdDashboard,
      path: `/${userName}/dashboard`,
    },
    {
      name: "Account",
      icon: MdAccountBox,
      path: `/${userName}/account`,
    },
    {
      name: "Results",
      icon: MdAssignment,
      path: `/${userName}/results`,
    },
    {
      name: "Schedule",
      icon: MdCalendarMonth,
      path: `/${userName}/schedule`,
    },
  ];
  return (
    <Box
      w="20vw"
      h="100%"
      m="20px"
      display="flex"
      flexDirection="column"
      position={"fixed"}
    >
      {sideMenuItems.map((item) => (
        <label key={item.name}>
          <Box mt="10px" ml="10px">
            <Button variant="ghost" onClick={() => navigate(item.path)}>
              <Icon as={item.icon} w={7} h={7} mr="13px" />
              {item.name}
            </Button>
          </Box>
        </label>
      ))}
    </Box>
  );
};

