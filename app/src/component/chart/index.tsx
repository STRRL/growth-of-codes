import { Line } from "react-chartjs-2";

import { Chart as ChartJS, registerables } from "chart.js";

ChartJS.register(...registerables);

export interface ChartProps {
  timeSeries: Point[];
}
export interface Point {
  time: Date;
  value: number;
}

function Chart(props: ChartProps) {
  const extractDataSeries = (timeSeries: Point[]) => {
    const data = timeSeries.map((point) => {
      return point.value;
    });
    const labels = timeSeries.map((point) => {
      return point.time;
    });
    return { labels, data };
  };
  const { labels, data } = extractDataSeries(props.timeSeries);

  return (
    <div>
      <Line
        data={{
          labels: labels,
          datasets: [
            {
              label: "random",
              data: data,
            },
          ],
        }}
      ></Line>
    </div>
  );
}

export default Chart;
