import json
import numpy as np
import matplotlib.pyplot as plt
from matplotlib.patches import Polygon
from pprint import pprint


class DataProcessor:
    def __init__(self, datafilename):
        with open(datafilename) as data_file:
            self.evaldata = json.load(data_file)
            self.speedups = {}
            self.php5_speedups = [1.0] * 10
            self.php7_speedups = []
            self.hhvm_php_speedups = []
            self.hhvm_speedups = [1.0] * 10
            self.python_speedups = []
            self.java_speedups = []
            self.go_speedups = []
            self.php_names = ["PHP 5.6", "PHP 7.0", "HHVM"]
            self.language_names = ["HHVM", "Python", "Java", "Go"]
            self.benchmark_name = {"fibonacci": "Fibonacci", "longarray": "Long Array", "nbody": "N-Body",
                                   "binarytress": "Binary Trees", "matrixmulti": "Matrix\nMultiplication",
                                   "mandelbrot": "Mandelbrot", "spectral-norm": "Spectral Norm",
                                   "parentheses": "Parentheses",
                                   "nqueens": "N Queens", "numofislands": "Num of\nIslands"}
            self.calc_speedups()

    def calc_speedups(self):
        for item, timespent in self.evaldata.iteritems():
            self.php7_speedups.append(timespent["php5.6"] / timespent["php7.0"])
            self.hhvm_php_speedups.append(timespent["php5.6"] / timespent["hhvm"])

            self.python_speedups.append(timespent["hhvm"] / timespent["python"])
            self.java_speedups.append(timespent["hhvm"] / timespent["java"])
            self.go_speedups.append(timespent["hhvm"] / timespent["go"])

    def plot_sppedups(self):
        php_data = [ self.php5_speedups, self.php7_speedups, self.hhvm_php_speedups]
        self.plot_barchart(php_data)
        alllang_data = [self.hhvm_speedups, self.python_speedups, self.java_speedups, self.go_speedups]
        self.plot_barchart(alllang_data)

    def plot_barchart(self, data_dict):
        n_groups = 10

        fig, ax = plt.subplots()
        index = np.arange(n_groups)
        if len(data_dict) == 3:
            bar_width = 0.25
        else:
            bar_width = 0.2

        opacity = 0.6
        error_config = {'ecolor': '0.3'}
        colors = ['r', 'g', 'b', 'm']
        i = 0
        maxValue = 0

        def autolabel(rects):
            """
            Attach a text label above each bar displaying its height
            """
            for rect in rects:
                height = rect.get_height()
                ax.text(rect.get_x() + rect.get_width() / 2., 1.05 * height,
                        '%.1f' % height,
                        ha='center', va='bottom')

        for data in data_dict:
            if len(data_dict) == 3:
                label = self.php_names[i]
            else:
                label = self.language_names[i]
            maxValue = max([maxValue, max(data)])
            rect = ax.bar(index + (i + 1) * bar_width, data, bar_width,
                          alpha=opacity,
                          color=colors[i],
                          error_kw=error_config,
                          label=label)
            autolabel(rect)
            i = i + 1

        plt.grid(True)
        ax.set_ylim(0, maxValue + 10)
        plt.xticks(index + bar_width * (len(data_dict) * 2-1) / 2, self.benchmark_name.values(), rotation=60)

        plt.xlabel('Benchmarks')
        plt.ylabel('Speedups')
        plt.title('Speedups by Benchmarks')


        plt.legend(loc="upper left")

        plt.tight_layout()
        plt.show()


dp = DataProcessor("BenchmarksData.json")
dp.plot_sppedups()
i = 3
