import { DefaultLayout } from "../components/DefaultLayout";
import { B3results } from "../components/B3results"

const B3ResultsPage = () => {

  interface problemStatus {
    id: number,
    Status: String,
  }

  interface B3Status {
    user: string,
    problems: problemStatus[],
  }

  interface B3StatusProps {
    data: B3Status[]
  }

  const b3StatusData: B3StatusProps = {
    data: [
      {
        user: 'isoda',
        problems: [
          { id: 1, Status: 'Complete' },
          { id: 2, Status: 'Incomplete' }
        ]
      },
      {
        user: 'araki',
        problems: [
          { id: 1, Status: 'Incomplete' },
          { id: 2, Status: 'Complete' }
        ]
      },
      {
        user: 'kimura',
        problems: [
          { id: 1, Status: 'Incomplete' },
          { id: 2, Status: 'Complete' }
        ]
      }
    ]
  };

  return (
    <DefaultLayout>
      <B3results data={b3StatusData.data} />
    </DefaultLayout>
  );
};

export default B3ResultsPage;
