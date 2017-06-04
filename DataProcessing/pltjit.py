import json
import numpy as np
import matplotlib.pyplot as plt
from matplotlib.patches import Polygon
from pprint import pprint

class DataProcessor:

    def __init__(self, datafilename):
        with open(datafilename) as data_file:
            self.evaldata = json.load(data_file)

    def plot_charts(self):
        fig, ax1 = plt.subplots()
        t = self.evaldata["n"]
        hhvm_time = self.evaldata["hhvm"]

        hhvm_single = [hhvm_time[i] * 1000000 / t[i] for i in range(len(t))]
        ax1.semilogx(t, hhvm_single, 'b^-', label='HHVM')
        ax1.set_xlabel('Num of runs')
        # Make the y-axis label, ticks and tick labels match the line color.
        ax1.set_ylabel('Time(us)')
        ax1.tick_params('y')

        php7_time = self.evaldata["php7.0"]
        php7_single = [php7_time[i] * 1000000 / t[i] for i in range(len(t))]
        ax1.semilogx(t, php7_single, 'r*-', label='PHP 7.0')

        php56_time = self.evaldata["php5.6"]
        php56_single = [php56_time[i] * 1000000 / t[i] for i in range(len(t))]
        ax1.semilogx(t, php56_single, 'go-', label='PHP 5.6')

        plt.title("Averaged function execution time")
        plt.xlim(400, 1.1e7)
        plt.legend()
        plt.grid(True)
        fig.tight_layout()
        plt.show()



dp = DataProcessor("jit.json")
dp.plot_charts()