import {
  LineChart,
  Line,
  XAxis,
  YAxis,
  Tooltip,
  CartesianGrid,
  ResponsiveContainer,
} from "recharts";
import moment from "moment";

export interface ChartProps {
  timeSeries: Point[];
}
export interface Point {
  time: Date;
  value: number;
}

function Chart(props: ChartProps) {
  return (
      <ResponsiveContainer>
        <LineChart
          data={props.timeSeries.map((point) => {
            return {
              time: moment(point.time).format("YYYY-MM-DD"),
              value: point.value,
            };
          })}
        >
          <CartesianGrid />
          <XAxis dataKey="time" />
          <YAxis />
          <Tooltip />
          <Line type="monotone" dataKey="value" stroke="#8884d8" />
        </LineChart>
      </ResponsiveContainer>
  );
}

export default Chart;
