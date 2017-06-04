import json
import numpy as np
import matplotlib.pyplot as plt
from matplotlib.patches import Polygon
from pprint import pprint

class DataProcessor:

    def __init__(self, datafilename):
        with open(datafilename) as data_file:
            self.evaldata = json.load(data_file)
            self.benchmark_name = {"mediawiki": "MediaWiki", "wordpress": "WordPress"}

    def plot_charts(self):
        for benchmark, tests in self.evaldata.iteritems():
            self.plot_benchmark_charts(benchmark, tests)
        plt.show()

    def plot_benchmark_charts(self, benchmark, testdata):
        speedups = []
        for tool, time in testdata.iteritems():
            randomDists = [ self.benchmark_name[benchmark] + "\n" + tool + "\nPHP 5.6",
                            self.benchmark_name[benchmark] + "\n" + tool + "\nPHP 7.0",
                            self.benchmark_name[benchmark] + "\n" + tool + "\nHHVM" ]
            data = [time["php5.6"], time['php7.0'], time['hhvm']]
            maxrange = max([max(data[0]), max(data[1]), max(data[2])]) + 10;
            minrange = min([min(data[0]), min(data[1]), min(data[2])]) - 10;

            timephp5 = np.average(data[0])
            timephp7 = np.average(data[1])
            timehhvm = np.average(data[2])

            speedups.append([1, timephp5/float(timephp7), timephp5/float(timehhvm)])

            fig, ax1 = plt.subplots(figsize=(10, 6))
            fig.canvas.set_window_title('A Boxplot Example')
            plt.subplots_adjust(left=0.075, right=0.95, top=0.9, bottom=0.25)

            bp = plt.boxplot(data, notch=0, sym='+', vert=1, whis=1.5)
            plt.setp(bp['boxes'], color='black')
            plt.setp(bp['whiskers'], color='black')
            plt.setp(bp['fliers'], color='red', marker='+')

            # Add a horizontal grid to the plot, but make it very light in color
            # so we can use it for reading data values but not be distracting
            ax1.yaxis.grid(True, linestyle='-', which='major', color='lightgrey', alpha=0.5)

            # Hide these grid behind plot objects
            ax1.set_axisbelow(True)
            ax1.set_title('Response Time of ' + self.benchmark_name[benchmark])
            ax1.set_ylabel('Response Time (ms)')

            # Now fill the boxes with desired colors
            boxColors = ['red', 'green','royalblue']
            numBoxes = len(data)
            medians = list(range(numBoxes))
            for i in range(numBoxes):
                box = bp['boxes'][i]
                boxX = []
                boxY = []
                for j in range(5):
                    boxX.append(box.get_xdata()[j])
                    boxY.append(box.get_ydata()[j])

                boxCoords = list(zip(boxX, boxY))
                # Alternate between Dark Khaki and Royal Blue
                k = i
                boxPolygon = Polygon(boxCoords, facecolor=boxColors[k])
                ax1.add_patch(boxPolygon)
                # Now draw the median lines back over what we just filled in
                med = bp['medians'][i]
                medianX = []
                medianY = []
                for j in range(2):
                    medianX.append(med.get_xdata()[j])
                    medianY.append(med.get_ydata()[j])
                    plt.plot(medianX, medianY, 'k')
                    medians[i] = medianY[0]
                # Finally, overplot the sample averages, with horizontal alignment
                # in the center of each box
                plt.plot([np.average(med.get_xdata())], [np.average(data[i])], color='w', marker='*', markeredgecolor='k')

            # Set the axes ranges and axes labels
            ax1.set_xlim(0.5, numBoxes + 0.5)
            top = maxrange
            bottom = minrange
            ax1.set_ylim(bottom, top)
            xtickNames = plt.setp(ax1, xticklabels=randomDists)
            plt.setp(xtickNames, rotation=45, fontsize=12)

            # Due to the Y-axis scale being different across samples, it can be
            # hard to compare differences in medians across the samples. Add upper
            # X-axis tick labels with the sample medians to aid in comparison
            # (just use two decimal places of precision)
            pos = np.arange(numBoxes) + 1
            upperLabels = [str(np.round(s, 2)) for s in medians]
            for tick, label in zip(range(numBoxes), ax1.get_xticklabels()):
                k = tick
                ax1.text(pos[tick], top - (top * 0.05), upperLabels[tick],
                     horizontalalignment='center', size='small', weight='semibold',
                     color=boxColors[k])

            # Finally, add a basic legend
            '''            
            plt.figtext(0.90, 0.12, 'PHP 5.6',
                        backgroundcolor=boxColors[0], color='white', weight='roman',
                        size='small')
            plt.figtext(0.90, 0.08, 'PHP 7.0',
                        backgroundcolor=boxColors[1],
                        color='white', weight='roman', size='small')
            plt.figtext(0.90, 0.04, 'HHVM 3.18',
                        backgroundcolor=boxColors[2],
                        color='white', weight='roman', size='small')
            
            '''

        speedups = np.asarray(speedups).T.tolist()

        ind = np.arange(len(speedups[0])) # the x locations for the groups
        width = 0.25  # the width of the bars

        fig, ax = plt.subplots()
        rects1 = ax.bar(width + ind, speedups[0], width, color='r')
        rects2 = ax.bar(ind + 2 * width, speedups[1], width, color='g')
        rects3 = ax.bar(ind + 3 * width, speedups[2], width, color='b')

        # add some text for labels, title and axes ticks
        ax.set_ylabel('Speedups')
        ax.set_title('Speedups of ' + self.benchmark_name[benchmark])
        ax.set_xticks(ind + width * 3 / 2 + width)
        ax.set_xticklabels(('Postman', 'CURL', 'WebpageTest'))
        ax.set_ylim(0, 1.5)

        ax.legend((rects1[0], rects2[0], rects3[0]), ('PHP 5.6', 'PHP 7.0', 'HHVM 3.18'), loc="upper left")

        #autolabel(rects1)
        #autolabel(rects2)

dp = DataProcessor("RawData.json")
dp.plot_charts()