import { Box, Button, Icon } from "@chakra-ui/react";
import {
  MdDashboard,
  MdAccountBox,
  MdAssignment,
  MdOutlineChecklist
} from "react-icons/md";
import { GiPodium } from "react-icons/gi";
import { useNavigate } from "react-router-dom";

export const SideMenu = () => {
  const navigate = useNavigate();
  const userName = localStorage.getItem("authUserName")
  const loginUserJoinedYear = new Date(localStorage.getItem("authJoinedDate") || Date.now()).getFullYear()
  const nowDate = new Date()
  const sideMenuItems = [
    {
      name: "Dashboard",
      icon: MdDashboard,
      path: `/dashboard`,
    },
    {
      name: "Account",
      icon: MdAccountBox,
      path: `/account`,
    },
    {
      name: "Results",
      icon: MdAssignment,
      path: `/results/${userName}`,
    },
    {
      name: "Petit Coder",
      icon: GiPodium,
      path: `/petit_coder`,
    },
    {
      name: "B3 Status",
      icon: MdOutlineChecklist,
      path: `/all_results`
    }
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
        item.name === "B3 Results" && loginUserJoinedYear === nowDate.getFullYear() ? null :
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

